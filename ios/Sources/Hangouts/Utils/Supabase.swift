//
//  Supabase.swift
//  Hangouts
//
//  Created by Stone Liu on 6/14/25.
//
import Supabase
import Foundation

// Create a new Supabase Client
public func createSupabaseClient() throws -> SupabaseClient {
    let envVariables = ProcessInfo.processInfo.environment
    guard let url = envVariables["SUPABASE_URL"] else {
        throw EnvironmentError.missingURL
    }

    guard let key = envVariables["SUPABASE_KEY"] else {
        throw EnvironmentError.missingKey
    }

    guard let parsedURL = URL(string: url) else {
        throw EnvironmentError.missingURL
    }

    let client = SupabaseClient(supabaseURL: parsedURL, supabaseKey: key)
    return client
}
