import Foundation
import OpenAPIRuntime
import OpenAPIURLSession

enum HangoutEnum<T> {
    case Response(T)
    case Error(String)
}

func HangoutWrapper<T>(fn: () async throws -> T) async -> HangoutEnum<T> {
    do {
        let res = try await fn()
        return HangoutEnum.Response(res)
    } catch {
        return HangoutEnum.Error("Unexpected error has occurred.")
    }
}

public class HangoutsService {
    var client: Client
    public init(url: URL) {
        self.client = Client(
            serverURL: url,
            transport: URLSessionTransport()
        )
    }

    func healthcheck() async -> HangoutEnum<Operations.GetHealthcheck.Output> {
        let fn = {
            try await self.client.getHealthcheck()
        }
        return await HangoutWrapper(fn: fn)
    }

    func createUser(firstName: String, lastName: String, username: String, age: Int, email: String)
        async
        -> HangoutEnum<
            Operations.PostApiV1User.Output
        >
    {
        let fn = {
            let jsonPayload = Operations.PostApiV1User.Input.Body.JsonPayload(
                firstName: firstName,
                lastName: lastName,
                username: username,
                email: email)

            let body = Operations.PostApiV1User.Input.Body.json(jsonPayload)
            let req = Operations.PostApiV1User.Input(body: body)
            return try await self.client.postApiV1User(req)
        }
        return await HangoutWrapper(fn: fn)
    }

}
