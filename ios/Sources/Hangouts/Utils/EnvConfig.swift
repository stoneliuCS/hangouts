import Foundation

enum EnvironmentError: Error {
    case missingURL
    case missingKey
    case invalidURL
}

// Holds flexible configurations for the client.
public struct EnvConfig {
    let API_BASE_URL: URL
    let SUPABASE_KEY: String
    let SUPABASE_URL: URL
}

// Creates an environment config based off the given environment variables.
public func createEnvConfig() -> EnvConfig {

    let apiBaseUrl =
        Bundle.main.object(forInfoDictionaryKey: "API_BASE_URL") as? String
    let supabaseKey =
        Bundle.main.object(forInfoDictionaryKey: "SUPABASE_KEY") as? String
    let supabaseUrl =
        Bundle.main.object(forInfoDictionaryKey: "SUPABASE_URL") as? String

    guard let apiBaseUrl = apiBaseUrl else {
        fatalError("No base api url found.")
    }

    guard let supabaseKey = supabaseKey else {
        fatalError("No Supabase api key found.")
    }

    guard let supabaseUrl = supabaseUrl else {
        fatalError("No Supabase url found.")
    }

    let apiBaseURL = URL(string: apiBaseUrl)
    let supabaseURL = URL(string: supabaseUrl)

    guard let apiBaseURL = apiBaseURL else {
        fatalError("Base api url is malformed.")
    }

    guard let supabaseURL = supabaseURL else {
        fatalError("Supabase url is malformed.")
    }

    return EnvConfig(
        API_BASE_URL: apiBaseURL, SUPABASE_KEY: supabaseKey,
        SUPABASE_URL: supabaseURL)

}
