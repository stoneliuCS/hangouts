import Foundation
import SwiftUI

@main
struct HangoutsApp: App {
    // Configurations for current environment
    private var userState: UserState = UserState(cfg: createEnvConfig())

    var body: some Scene {
        WindowGroup {
            NavigationStack {
                if !userState.isLoggedIn() {
                    LandingView()
                } else {
                    HomeView()
                }

            }.environmentObject(userState)
        }
    }
}
