import SwiftUI

struct LandingView: View {

    @State private var opacity = 0.0
    @State private var paddingVertical = 100.0
    @State private var easeInDuration = 2.0

    var body: some View {
        VStack(spacing: 150) {
            VStack {
                Text("Hangouts.ai").font(.system(size: 60)).italic().bold().opacity(
                    opacity
                )
                .animation(.easeInOut(duration: 2.0), value: opacity).onAppear { opacity = 1.0 }
                Text("Plan less, do more.")
                    .font(.system(size: 24))
                    .foregroundColor(.secondary)
                    .opacity(opacity)
                    .animation(.easeInOut(duration: 2.0).delay(0.5), value: opacity)
            }
            // Buttons
            VStack(spacing: 30) {
                Button(action: {
                }) {
                    Text("Log In")
                        .font(.title).bold()
                        .underline()
                        .foregroundColor(Color.black)
                }
                Button(action: {
                }) {
                    Text("Sign Up")
                        .font(.title).bold()
                        .underline()
                        .foregroundColor(Color.black)
                }

            }
            .opacity(opacity)
            .animation(.easeInOut(duration: 2.0).delay(1.0), value: opacity)
        }.padding(.bottom, paddingVertical)
    }

}
