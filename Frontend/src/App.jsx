import "./App.css";
import Home from "./Pages/Home/Home";
import Temp from "./Pages/Temp/Temp";
import {
	BrowserRouter as Router,
	Route,
	Routes,
	Navigate,
} from "react-router-dom";
import CreateBook from "./Pages/CreateBook/CreateBook";
import AdminDashboard from "./Pages/Dashboard/AdminDashboard";
import Login from "./Pages/Login_Signup/Login";
import Signup from "./Pages/Login_Signup/Signup";
import CreateLibrary from "./Pages/CreateLibrary/CreateLibrary";
import Profile from "./Pages/Profile/Profile";
import ProtectedRoute from "./ProtectedRoute";
import ManageBooks from "./Pages/ManageBooks/ManageBooks";

function App() {
	return (
		<>
			<div>
				<Router>
					<Routes>
						<Route path="/" element={<Home />} />
						<Route path="/temp" element={<Temp />} />
						<Route path="/signup" element={<Signup />} />
						<Route path="/login" element={<Login />} />
						<Route
							path="/profile"
							element={<ProtectedRoute element={<Profile />} />}
						/>
						<Route
							path="/dashboard"
							element={
								<ProtectedRoute element={<AdminDashboard />} />
							}
						/>
						<Route
							path="/create-library"
							element={
								<ProtectedRoute element={<CreateLibrary />} />
							}
						/>
						<Route
							path="/create-book"
							element={
								<ProtectedRoute element={<CreateBook />} />
							}
						/>
						<Route
							path="/manage-books-reader"
							element={
								<ProtectedRoute element={<ManageBooks />} />
							}
						/>
						<Route path="*" element={<Navigate to="/" replace />} />
					</Routes>
				</Router>
			</div>
		</>
	);
}

export default App;
