import Foundation
import Supabase

struct ErrorRes {
    var message: String
}

// Root State Management
@MainActor
public class UserState: ObservableObject {

    @Published var isAuthenticated: Bool
    private var session: Session?
    private var authService: AuthService
    private var userId: String?
    private var email: String?
    private var onboarded: Bool
    private let localStore: NSObject

    init(cfg: EnvConfig) {
        self.authService = AuthService(supabaseURL: cfg.SUPABASE_URL, supabaseKey: cfg.SUPABASE_KEY)
        self.onboarded = false
        self.localStore = UserDefaults.standard
        self.isAuthenticated = false
    }

    // Registers the user by email, returning an error response if failed to do so.
    func registerByEmail(email: String, password: String) async -> ErrorRes? {
        let res = await self.authService.registerWithEmail(email: email, password: password)

        guard let res = res else {
            return ErrorRes(message: "Failed to register with email and password.")
        }
        self.email = email
        self.session = res.session
        self.isAuthenticated = true
        return nil
    }

}
