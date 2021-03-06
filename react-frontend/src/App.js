import { ClerkProvider, SignedIn, SignedOut, SignIn } from "@clerk/clerk-react";
import Home from "./pages/Home/Home";
import "./css/styles.css";

function App() {
  // Retrieve Clerk settings from the environment
  const clerkFrontendApi = process.env.REACT_APP_CLERK_FRONTEND_API;
  return (
    <ClerkProvider frontendApi={clerkFrontendApi}>
      <div className="app-background">
        <SignedOut>
          <SignIn />
        </SignedOut>
        <SignedIn>
          <div className="app-container">
            <Home />
          </div>
        </SignedIn>
      </div>
    </ClerkProvider>
  );
}

export default App;
