import Foundation
import Supabase

struct UserStore {
    // Current UUID for the user or none if the user hasn't signed in or registered.
    var userId: String?
    // Current email for the user or none if the user hasn't registed or used an email to sign up.
    var email: String?
    // A Boolean flag to check if the user has viewed the onboarding screens.
    var onboarded: Bool
    // Current session information
    var session: Session?
}

struct ErrorRes {
    var message: String
}

// Root State Management
public class UserState: ObservableObject {

    private var authService: AuthService
    private var userStore: UserStore

    init(cfg: EnvConfig) {
        self.authService = AuthService(supabaseURL: cfg.SUPABASE_URL, supabaseKey: cfg.SUPABASE_KEY)
        self.userStore = UserStore(userId: nil, email: nil, onboarded: true, session: nil)
    }

    func isLoggedIn() -> Bool {
        guard let userSession = self.userStore.session else {
            return false
        }
        return userSession.isExpired
    }

    // Registers the user by email, returning an error response if failed to do so.
    func registerByEmail(email: String, password: String) async -> ErrorRes? {
        let res = await self.authService.registerWithEmail(email: email, password: password)

        guard let res = res else {
            return ErrorRes(message: "Failed to register with email and password.")
        }
        // Assign values to the user store.
        self.userStore.email = email
        self.userStore.session = res.session
        return nil
    }

}
