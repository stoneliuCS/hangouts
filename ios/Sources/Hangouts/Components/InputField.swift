import SwiftUI

enum InputType {
    case text
    case secure
}

struct InputField: View {
    let placeholder: String
    @Binding var text: String
    var inputType: InputType

    var body: some View {
        Group {
            switch self.inputType {
            case .text:
                TextField(placeholder, text: $text)
            case .secure:
                SecureField(placeholder, text: $text)
            }
        }
        .padding()
        .cornerRadius(8)
        .overlay(
            RoundedRectangle(cornerRadius: 8)
                .stroke(Color.black, lineWidth: 1)
        ).textInputAutocapitalization(.never).disableAutocorrection(true)
    }
}
