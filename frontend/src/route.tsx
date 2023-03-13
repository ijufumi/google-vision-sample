import React, { FC } from "react"
import { BrowserRouter, Routes, Route } from "react-router-dom";
import App from "./pages/App"
import ResultPage from "./pages/ResultPage"

export interface Props {}

const AppRoute: FC <Props> = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route
          path="/:jobId"
          element={<ResultPage />}
        />
        <Route
          path="*"
          element={<App />}
        />
      </Routes>
    </BrowserRouter>
  )
}

export default AppRoute
