import { useNavigate } from "react-router-dom";

export const logout = () => {
    const navigate = useNavigate();
    localStorage.removeItem("token");
    localStorage.removeItem("user_id");
    navigate("/login");
    console.log("User logged-out successfully !!");
};