import Foundation
import Supabase

enum AuthEnum<T> {
    case Response(T)
    case Error(String)
}

func AuthWrapper<T>(fn: () async throws -> T) async -> AuthEnum<T> {
    do {
        let res = try await fn()
        return AuthEnum.Response(res)
    } catch let error as Supabase.AuthError {
        return AuthEnum.Error(error.message)
    } catch {
        return AuthEnum.Error("Unexpected error has occurred.")
    }
}

// Handles client authorization with Supabase being the client server.
public class AuthService {

    private var supabaseClient: SupabaseClient

    init(supabaseURL: URL, supabaseKey: String) {
        self.supabaseClient = SupabaseClient(supabaseURL: supabaseURL, supabaseKey: supabaseKey)
    }

    // Registers the user into supabase
    func registerWithEmail(email: String, password: String) async -> AuthEnum<AuthResponse> {
        let fn = {
            try await self.supabaseClient.auth.signUp(
                email: email,
                password: password
            )
        }
        return await AuthWrapper(fn: fn)
    }

    // Logs in with Email
    func loginWithEmail(email: String, password: String) async -> AuthEnum<Session> {
        let fn = {
            try await self.supabaseClient.auth.signIn(email: email, password: password)
        }
        return await AuthWrapper(fn: fn)
    }

}
