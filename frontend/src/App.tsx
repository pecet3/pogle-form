import "./App.css";
import { Route, Routes } from "react-router-dom";
import { ProtectedPage } from "./wrappers/Protected";
import { Home } from "./pages/Home";

import { Auth } from "./pages/Auth";
import { Courses } from "./pages/Courses";
import { Persons } from "./pages/Persons";

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
          path="/persons"
          element={
            <ProtectedPage>
              <Persons />
            </ProtectedPage>
          }
        />
        <Route
          path="/courses"
          element={
            <ProtectedPage>
              <Courses />
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
