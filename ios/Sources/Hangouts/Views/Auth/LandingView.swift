import SwiftUI

struct LandingView: View {

    @State private var opacity = 0.0
    @State private var paddingVertical = 100.0
    @State private var easeInDuration = 2.0
    @State private var path = NavigationPath()

    var body: some View {
        NavigationStack(path: $path) {
            VStack {
                Image("logo").resizable()
                VStack(spacing: 30) {
                    Button("Log In") {
                        path.append("Login")
                    }.font(.title2)
                    Button("Sign Up") {
                        path.append("SignUp")
                    }.font(.title2)
                }
            }
            .opacity(opacity)
            .animation(.easeInOut(duration: 2.0).delay(1.0), value: opacity)
            .navigationTitle("Hangouts.ai")
            .onAppear {
                opacity = 1.0
            }
            .navigationDestination(for: String.self) { value in
                if value == "Login" {
                    LoginView()
                } else if value == "SignUp" {
                    SignUpView()
                }
            }.padding(.bottom, paddingVertical)

        }
    }
}
