import React from "react";
import { Navigate } from "react-router-dom";

const ProtectedRoute = ({ element: Component }) => {
    const isAuthenticated = !!localStorage.getItem("token"); // Example authentication check

    return isAuthenticated ? Component : <Navigate to="/login" replace />;
};

export default ProtectedRoute;