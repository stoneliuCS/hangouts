import Foundation
import OpenAPIRuntime
import OpenAPIURLSession

public class HangoutsService {
    var client: Client
    public init(url: URL) {
        self.client = Client(
            serverURL: url,
            transport: URLSessionTransport()
        )
    }

    func healthcheck() async throws -> Operations.GetHealthcheck.Output {
        return try await self.client.getHealthcheck()
    }
}
