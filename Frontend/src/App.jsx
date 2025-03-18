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
import AdminDashboard from "./Pages/AdminDashboard/AdminDashboard";
import Login from "./Pages/Login_Signup/Login";
import Signup from "./Pages/Login_Signup/Signup";
import CreateLibrary from "./Pages/CreateLibrary/CreateLibrary";
import Profile from "./Pages/Profile/Profile";

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
						<Route path="/profile" element={<Profile />} />
						<Route path="/create-book" element={<CreateBook />} />
						<Route path="/dashboard" element={<AdminDashboard />} />
						<Route
							path="/create-library"
							element={<CreateLibrary />}
						/>
						<Route path="*" element={<Navigate to="/" replace />} />
					</Routes>
				</Router>
			</div>
		</>
	);
}

export default App;
