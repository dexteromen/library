import React, { useState } from "react";
import "./Signin.css";

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
			<div className="wrapper-signin">
				<form className="form" onSubmit={handleSubmit}>
					<div className="field-div">
						<label>Email:</label>
						<input
							type="email"
							value={email}
							onChange={(e) => setEmail(e.target.value)}
						/>
						{errors.email && <p>{errors.email}</p>}
					</div>
					<div className="field-div">
						<label>Password:</label>
						<input
							type="password"
							value={password}
							onChange={(e) => setPassword(e.target.value)}
						/>
						{errors.password && <p>{errors.password}</p>}
					</div>
					<button type="submit">Sign In</button>
				</form>
			</div>
		</>
	);
}

export default Signin;
