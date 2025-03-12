import "./App.css";
import Home from "./Pages/Home/Home";
import Signin from "./Pages/Signin/Signin";
import Signup from "./Pages/Signup/Signup";

function App() {
	return (
		<>
			<div>
				{/* <Home /> */}
				<Signup />
				<Signin />
			</div>
		</>
	);
}

export default App;
