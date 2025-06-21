import Foundation
import SwiftUI

@main
struct HangoutsApp: App {
    // Configurations for current environment
    @StateObject private var userState: UserState = UserState(cfg: createEnvConfig())

    var body: some Scene {
        WindowGroup {
            Group {
                if !userState.isAuthenticated {
                    LandingView()
                } else {
                    HomeView()
                }
            }.environmentObject(userState)
        }
    }
}
