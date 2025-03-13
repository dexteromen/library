import React, { useState } from "react";
import { Link } from "react-router-dom";
import "./Auth.css";
import Navbar from "../../Components/Navbar/Navbar";

function Login() {
	const [formData, setFormData] = useState({
		email: "",
		password: "",
	});

	const [errors, setErrors] = useState({});

	const handleChange = (e) => {
		const { name, value } = e.target;
		setFormData({
			...formData,
			[name]: value,
		});
	};

	const validate = () => {
		let tempErrors = {};

		if (!formData.email) tempErrors.email = "Email is required";
		if (!/\S+@\S+\.\S+/.test(formData.email))
			tempErrors.email = "Invalid email";
		if (!formData.password) tempErrors.password = "Password is required";
		if (formData.password.length < 6)
			tempErrors.password = "Password must be at least 6 characters";

		setErrors(tempErrors);
		return Object.keys(tempErrors).length === 0;
	};

	const handleSubmit = (e) => {
		e.preventDefault();
		if (validate()) {
			//Submit details
			console.log("Login Data Sent Successfully.");
			console.log(formData);
		}
	};

	return (
		<>
			<Navbar />
			<div className="form-wrapper">
				<div className="login_form">
					<form action="#" onSubmit={handleSubmit}>
						<h1>LOGIN</h1>
						<div className="input_box">
							<label>Email</label>
							<input
								type="email"
								name="email"
								value={formData.email}
								placeholder="Enter email address"
								onChange={handleChange}
								required
							/>
							{errors.email && <p>{errors.email}</p>}
						</div>
						<div className="input_box">
							<div className="password_title">
								<label>Password</label>
								{/* <a href="#">Forgot Password?</a> */}
							</div>
							<input
								type="password"
								name="password"
								value={formData.password}
								placeholder="Enter your password"
								onChange={handleChange}
								required
							/>
							{errors.password && <p>{errors.password}</p>}
						</div>
						<button type="submit">LOGIN</button>
						<p className="sign_up">
							Don't have an account?{" "}
							<Link to="/signup">SIGNUP</Link>
						</p>
					</form>
				</div>
			</div>
		</>
	);
}

export default Login;
