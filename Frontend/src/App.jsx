import "./App.css";
import Home from "./Pages/Home/Home";
import Signin from "./Pages/Signin/Signin";
import Signup from "./Pages/Signup/Signup";
import Temp from "./Pages/Temp";
import {
	BrowserRouter as Router,
	Route,
	Routes,
	Navigate,
} from "react-router-dom";

function App() {
	return (
		<>
			<div>
				<Router>
					<Routes>
						<Route path="/" element={<Home />} />
						<Route path="/signup" element={<Signup />} />
						<Route path="/signin" element={<Signin />} />
						<Route path="/temp" element={<Temp />} />
						<Route path="*" element={<Navigate to="/" replace />} />
					</Routes>
				</Router>
			</div>
		</>
	);
}

export default App;
