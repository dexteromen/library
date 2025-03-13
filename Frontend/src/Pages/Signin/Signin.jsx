import React, { useState } from "react";
import "./Signin.css";
import Navbar from "../../Components/Navbar/Navbar";

function Signin() {
	const [email, setEmail] = useState("");
	const [password, setPassword] = useState("");
	const [errors, setErrors] = useState({});

	const validate = () => {
		let errors = {};
		if (!email) {
			errors.email = "Email is required";
		} else if (!/\S+@\S+\.\S+/.test(email)) {
			errors.email = "Email address is invalid";
		}
		if (!password) {
			errors.password = "Password is required";
		} else if (password.length < 6) {
			errors.password = "Password must be at least 6 characters";
		}
		return errors;
	};

	const handleSubmit = (e) => {
		e.preventDefault();
		const validationErrors = validate();
		if (Object.keys(validationErrors).length === 0) {
			// Submit the form
			console.log("Form submitted");
			console.log("Email:" + { email });
			console.log("Password:" + { password });
		} else {
			setErrors(validationErrors);
		}
	};

	return (
		<>
			<Navbar />
			<div className="form-wrapper">
				<div className="login_form">
					<form action="#" onSubmit={handleSubmit}>
						<h1>LOGIN</h1>
						{/* <div className="login_option">
						<div className="option">
							<a href="#">
								<img src="/google.png" alt="Google" />
								<span>Google</span>
							</a>
						</div>
						<div className="option">
							<a href="#">
								<img src="/apple.png" alt="Apple" />
								<span>Apple</span>
							</a>
						</div>
					</div> */}
						{/* <p className="separator">
						<span>or</span>
					</p> */}
						<div className="input_box">
							<label>Email</label>
							<input
								type="email"
								id="email"
								value={email}
								placeholder="Enter email address"
								onChange={(e) => setEmail(e.target.value)}
								required
							/>
							{errors.email && <p>{errors.email}</p>}
						</div>
						<div className="input_box">
							<div className="password_title">
								<label>Password</label>
								<a href="#">Forgot Password?</a>
							</div>
							<input
								type="password"
								id="password"
								value={password}
								placeholder="Enter your password"
								onChange={(e) => setPassword(e.target.value)}
								required
							/>
							{errors.password && <p>{errors.password}</p>}
						</div>
						<button type="submit">Log In</button>
						<p className="sign_up">
							Don't have an account? <a href="#">Sign up</a>
						</p>
					</form>
				</div>
			</div>
		</>
	);
}

export default Signin;
