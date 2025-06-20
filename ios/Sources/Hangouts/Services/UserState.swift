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
    private var onboarded: Bool
    private let localStore: NSObject

    init(cfg: EnvConfig) {
        self.authService = AuthService(supabaseURL: cfg.SUPABASE_URL, supabaseKey: cfg.SUPABASE_KEY)
        self.onboarded = false
        self.localStore = UserDefaults.standard
        self.isAuthenticated = false
    }

    private func setAuth(session: Session?) {
        self.session = session
        self.isAuthenticated = true
    }

    // Registers the user by email, returning an error response if failed to do so.
    func registerByEmail(email: String, password: String) async -> ErrorRes? {
        let res = await self.authService.registerWithEmail(email: email, password: password)
        switch res {
        case .Response(let response):
            self.setAuth(session: response.session)
            return nil
        case .Error(let error):
            return ErrorRes(message: error)
        }
    }

    func signupByEmail(email: String, password: String) async -> ErrorRes? {
        let res = await self.authService.loginWithEmail(email: email, password: password)
        switch res {
        case .Response(let response):
            self.setAuth(session: response)
            return nil
        case .Error(let error):
            return ErrorRes(message: error)
        }
    }

    func logout() async -> ErrorRes? {
        let res = await self.authService.logout()
        switch res {
            case .Response():
                self.isAuthenticated = false
                return nil
            case .Error(let error):
                return ErrorRes(message : error)
        }
    }

}
