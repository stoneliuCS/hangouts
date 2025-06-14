import OpenAPIRuntime
import OpenAPIURLSession
import Foundation

enum EnvironmentError: Error {
    case missingURL
    case missingKey
}

public struct HangoutsClient {
    var client: Client
    public init(url : URL) {
        self.client = Client(
            serverURL: url,
            transport: URLSessionTransport()
        )
    }

    func healthcheck() async throws -> Operations.GetHealthcheck.Output {
        return try await self.client.getHealthcheck()
    }
}

