//
//  Auth.swift
//  Hangouts
//
//  Created by Stone Liu on 6/14/25.
//
import Supabase

enum AuthError : Error {
    case SignUpError
}

// Handles client authorization with Supabase being the client server.
class AuthService {

    private var supabaseClient: SupabaseClient

    init(supabaseClient: SupabaseClient) {
        self.supabaseClient = supabaseClient
    }

    // Registers the user into supabase
    func register(email: String, password: String) async throws -> Session {
        let res = try await self.supabaseClient.auth.signUp(
            email: email,
            password: password
        )
        guard res.session != nil else {
            throw AuthError.SignUpError
        }
        return res.session!
    }
    
}
