import FormValidator
import SwiftUI

class LoginForm: ObservableObject {
    @Published
    var manager = FormManager(validationType: .immediate)
    @FormField(validator: EmailValidator(message: "Email cannot be non empty."))
    var email: String = ""
    @FormField(validator: NonEmptyValidator(message: "Password cannot be non empty."))
    var password: String = ""
    lazy var emailValidation = _email.validation(manager: manager)
    lazy var passwordValidation = _password.validation(manager: manager)
}

struct LoginView: View {

    @ObservedObject private var signupForm = LoginForm()
    @EnvironmentObject private var userState: UserState
    @State private var isLoading = false
    @State private var errorMessage: String?
    @Environment(\.dismiss) private var dismiss
    var body: some View {

        VStack(spacing: 15) {
            InputField(placeholder: "Email", text: $signupForm.email, inputType: .text)
                .validation(signupForm.emailValidation).padding(.horizontal)
            InputField(placeholder: "Password", text: $signupForm.password, inputType: .secure)
                .validation(signupForm.passwordValidation).padding(.horizontal)

            if let errorMessage = errorMessage {
                Text(errorMessage).foregroundColor(.red).padding().background(
                    Color.red.opacity(0.1)
                ).cornerRadius(6)
            }

            Button("Log in!", action: { Task { await self.submit() } }).disabled(isLoading)
        }
    }

    private func submit() async {
        let valid = signupForm.manager.triggerValidation()
        if !valid {
            self.errorMessage = "Login submission incomplete, please check validation for errors."
            return
        }
        self.isLoading = true
        let res = await self.userState.signupByEmail(
            email: self.signupForm.email, password: self.signupForm.password)
        if let res = res {
            self.errorMessage = res.message
            return
        }
        self.isLoading = false
        self.dismiss()
    }
}
