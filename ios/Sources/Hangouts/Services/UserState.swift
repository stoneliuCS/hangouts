import Foundation
import Supabase

struct ErrorRes {
    var message: String
}

// Root User State Management
@MainActor
public class UserState: ObservableObject {

    @Published private var session: Session?
    private let authService: AuthService
    private let hangoutsService: HangoutsService
    private let localStore: NSObject

    public var isAuthenticated: Bool {
        session != nil && !session!.isExpired
    }
    @Published public var onboarded: Bool

    init(cfg: EnvConfig) {
        self.session = nil
        self.authService = AuthService(supabaseURL: cfg.SUPABASE_URL, supabaseKey: cfg.SUPABASE_KEY)
        self.hangoutsService = HangoutsService(url: cfg.API_BASE_URL)
        self.onboarded = false
        self.localStore = UserDefaults.standard
    }

    func registerUser(username: String, firstName: String, lastName: String, email: String) {

    }

    // Registers the user by email, returning an error response if failed to do so.
    func registerByEmail(email: String, password: String) async -> ErrorRes? {
        let res = await self.authService.registerWithEmail(email: email, password: password)
        switch res {
        case .Response(let response):
            self.session = response.session
            return nil
        case .Error(let error):
            return ErrorRes(message: error)
        }
    }

    func signupByEmail(email: String, password: String) async -> ErrorRes? {
        let res = await self.authService.loginWithEmail(email: email, password: password)
        switch res {
        case .Response(let response):
            self.session = response
            return nil
        case .Error(let error):
            return ErrorRes(message: error)
        }
    }

    func logout() async -> ErrorRes? {
        let res = await self.authService.logout()
        switch res {
        case .Response():
            self.session = nil
            return nil
        case .Error(let error):
            return ErrorRes(message: error)
        }
    }

}
