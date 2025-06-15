import Foundation
import Supabase

enum AuthError: Error {
    case SignUpError
}

// Handles client authorization with Supabase being the client server.
public class AuthService {

    private var supabaseClient: SupabaseClient

    init(supabaseURL: URL, supabaseKey: String) {
        self.supabaseClient = SupabaseClient(supabaseURL: supabaseURL, supabaseKey: supabaseKey)
    }

    // Registers the user into supabase
    func registerWithEmail(email: String, password: String) async -> AuthResponse? {
        return try? await self.supabaseClient.auth.signUp(
            email: email,
            password: password
        )
    }

    // Logs in with Email
    func loginWithEmail(email: String, password: String) async -> Session? {
        return try? await self.supabaseClient.auth.signIn(email: email, password: password)
    }

}
