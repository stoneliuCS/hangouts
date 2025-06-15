import Foundation
import Supabase

// Root State Management
public class UserState: ObservableObject {

    private var cfg: EnvConfig
    private var authService: AuthService
    private var userSession: Session?

    init(cfg: EnvConfig) {
        self.cfg = cfg
        self.authService = AuthService(supabaseURL: cfg.SUPABASE_URL, supabaseKey: cfg.SUPABASE_KEY)
    }

    func isLoggedIn() -> Bool {
        guard let userSession = self.userSession else {
            return false
        }
        return userSession.isExpired
    }

}
