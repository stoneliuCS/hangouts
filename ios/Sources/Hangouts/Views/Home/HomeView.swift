import SwiftUI

struct HomeView: View {
    @EnvironmentObject private var userState: UserState
    @State private var isLoading = false
    @State private var errorMessage: String?
    @Environment(\.dismiss) private var dismiss
    var body: some View {
        Text("HOME")
        Button("Sign out", action: { Task { await self.submit() } }).disabled(isLoading)
    }

    private func submit() async {
        self.isLoading = true
        let res = await self.userState.logout()
        if let res = res {
            self.isLoading = false
            self.errorMessage = res.message
            return
        }
        self.isLoading = false
        self.dismiss()
    }
}
