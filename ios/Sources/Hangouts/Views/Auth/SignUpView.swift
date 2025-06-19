import FormValidator
import SwiftUI

class SignupForm: ObservableObject {
    private static let passwordRegex = try! NSRegularExpression(
        pattern: "^(?=.*[A-Za-z])(?=.*\\d)[A-Za-z\\d]{8,}$", options: [])
    @Published
    var manager = FormManager(validationType: .immediate)

    @FormField(validator: NonEmptyValidator(message: "Username is required!"))
    var username: String = ""

    @FormField(validator: NonEmptyValidator(message: "First Name is required!"))
    var firstName: String = ""

    @FormField(validator: NonEmptyValidator(message: "Last Name is required!"))
    var lastName: String = ""

    @FormField(validator: EmailValidator(message: "Valid email is required!"))
    var email: String = ""

    @FormField(
        validator: PasswordValidator(
            pattern: passwordRegex,
            message:
                "Passwords must be atleast 8 characters long, contain at least one digit and only be alphanumeric"
        ))
    var password: String = ""

    lazy var usernameValidation = _username.validation(manager: manager)
    lazy var firstNameValidation = _firstName.validation(manager: manager)
    lazy var lastNameValidation = _lastName.validation(manager: manager)
    lazy var emailValidation = _email.validation(manager: manager)
    lazy var passwordValidation = _password.validation(manager: manager)
}

struct SignUpView: View {
    @ObservedObject private var signupForm = SignupForm()
    @EnvironmentObject private var userState: UserState
    @State private var isLoading = false
    @State private var errorMessage: String?
    @Environment(\.dismiss) private var dismiss
    var body: some View {
        VStack(spacing: 15) {
            InputField(placeholder: "Username", text: $signupForm.username, inputType: .text)
                .validation(signupForm.usernameValidation).padding(.horizontal)
            InputField(placeholder: "First Name", text: $signupForm.firstName, inputType: .text)
                .validation(signupForm.firstNameValidation).padding(.horizontal)
            InputField(placeholder: "Last Name", text: $signupForm.lastName, inputType: .text)
                .validation(signupForm.lastNameValidation).padding(.horizontal)
            InputField(placeholder: "Email", text: $signupForm.email, inputType: .text)
                .validation(signupForm.emailValidation).padding(.horizontal)
            InputField(placeholder: "Password", text: $signupForm.password, inputType: .secure)
                .validation(signupForm.passwordValidation).padding(.horizontal)

            if let errorMessage = errorMessage {
                Text(errorMessage).foregroundColor(.red).padding().background(
                    Color.red.opacity(0.1)
                ).cornerRadius(6)
            }

            Button("Let's Hang Out!", action: { Task { await self.submit() } }).disabled(isLoading)
        }
    }
    private func submit() async {
        let valid = signupForm.manager.triggerValidation()
        if !valid {
            self.errorMessage = "Form submission incomplete, please check for errors."
            return
        }
        self.isLoading = true
        let res = await self.userState.registerByEmail(
            email: self.signupForm.email, password: self.signupForm.password)
        if let res = res {
            self.errorMessage = res.message
        }
        self.isLoading = false
        self.dismiss()
    }
}
