import "./App.css";
import Login from "./Pages/Auth/Login";
import Home from "./Pages/Home/Home";
import Signup from "./Pages/Auth/Signup";
import Temp from "./Pages/Temp/Temp";
import {
	BrowserRouter as Router,
	Route,
	Routes,
	Navigate,
} from "react-router-dom";
import CreateBook from "./Pages/CreateBook/CreateBook";
import AdminDashboard from "./Pages/AdminDashboard/AdminDashboard";

function App() {
	return (
		<>
			<div>
				<Router>
					<Routes>
						<Route path="/" element={<Home />} />
						<Route path="/signup" element={<Signup />} />
						<Route path="/login" element={<Login />} />
						<Route path="/temp" element={<Temp />} />
						<Route path="/create-book" element={<CreateBook />} />
						<Route path="/dashboard" element={<AdminDashboard />} />
						<Route path="*" element={<Navigate to="/" replace />} />
					</Routes>
				</Router>
			</div>
		</>
	);
}

export default App;
