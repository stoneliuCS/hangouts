import Foundation
import Supabase

struct ErrorRes {
    var message: String
}

// Root State Management
@MainActor
public class UserState: ObservableObject {

    @Published private var session: Session?
    private var authService: AuthService
    @Published private var userId: String?
    @Published private var email: String?
    @Published private var onboarded: Bool

    init(cfg: EnvConfig) {
        self.authService = AuthService(supabaseURL: cfg.SUPABASE_URL, supabaseKey: cfg.SUPABASE_KEY)
        self.onboarded = false
    }

    public var isLoggedIn: Bool {
        guard let userSession = self.session else {
            return false
        }
        return !userSession.isExpired
    }

    // Registers the user by email, returning an error response if failed to do so.
    func registerByEmail(email: String, password: String) async -> ErrorRes? {
        let res = await self.authService.registerWithEmail(email: email, password: password)

        guard let res = res else {
            return ErrorRes(message: "Failed to register with email and password.")
        }
        // Assign values to the user store.
        self.email = email
        self.session = res.session
        return nil
    }

}
