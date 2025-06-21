import SwiftUI

struct HomeView: View {
    @EnvironmentObject private var userState: UserState
    @State private var isLoading = false
    @State private var errorMessage: String?
    @Environment(\.dismiss) private var dismiss
    var body: some View {
        TabView {
            HangoutsView()
                .tabItem {
                    Image(systemName: "person.2.fill")
                    Text("Hangouts")
                }
            IdeateView()
                .tabItem {
                    Image(systemName: "lightbulb.fill")
                    Text("Ideate")
                }
            ProfileView()
                .tabItem {
                    Image(systemName: "person.circle.fill")
                    Text("Profile")
                }
        }
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

struct HangoutsView: View {
    var body: some View {
        Text("I am Hangouts")
    }
}

struct IdeateView: View {
    var body: some View {
        Text("I am Ideate")
    }
}

struct ProfileView: View {
    var body: some View {

        Text("I am profile")
    }
}
