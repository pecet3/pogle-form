import "./App.css";
import { Route, Routes } from "react-router-dom";
import { ProtectedPage } from "./wrappers/Protected";
import { Home } from "./pages/Home";

import { Auth } from "./pages/Auth";

function App() {
  return (
    <>
      <Routes>
        <Route
          path="/"
          element={
            <ProtectedPage>
              <Home />
            </ProtectedPage>
          }
        />
        <Route
          path="/users"
          element={
            <ProtectedPage>
              <Home />
            </ProtectedPage>
          }
        />
        <Route
          path="/courses"
          element={
            <ProtectedPage>
              <Home />
            </ProtectedPage>
          }
        />
        <Route
          path="/auth"
          element={
            <>
              <Auth />
            </>
          }
        />
      </Routes>
    </>
  );
}

export default App;
