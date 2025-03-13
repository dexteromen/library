import React, { useState } from "react";
import { Link } from "react-router-dom";
import "./Auth.css";
import Navbar from "../../Components/Navbar/Navbar";

function Signup() {
	const [formData, setFormData] = useState({
		// username: "",
		name: "",
		contact_number: "",
		email: "",
		password: "",
		confirmPassword: "",
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
		// if (!formData.username) tempErrors.username = "Username is required";
		if (!formData.name) tempErrors.name = "Name is required";
		if (formData.contact_number.length != 10)
			tempErrors.contact_number =
				"Contact number is required only 10 digits";
		if (!formData.email) tempErrors.email = "Email is required";
		if (!/\S+@\S+\.\S+/.test(formData.email))
			tempErrors.email = "Email is invalid";
		if (!formData.password) tempErrors.password = "Password is required";
		if (formData.password.length < 6)
			tempErrors.password = "Password must be at least 6 characters";
		if (formData.password !== formData.confirmPassword)
			tempErrors.confirmPassword = "Passwords do not match";
		setErrors(tempErrors);
		return Object.keys(tempErrors).length === 0;
	};

	const handleSubmit = (e) => {
		e.preventDefault();
		if (validate()) {
			// Submit form
			console.log("Form submitted successfully");
			console.log(formData);
		}
	};
	return (
		<>
			<Navbar />
			<div className="form-wrapper">
				<div className="login_form">
					<form className="form" onSubmit={handleSubmit}>
						<h1>SIGNUP</h1>
						<div className="input_box">
							<label>Name</label>
							<input
								type="text"
								name="name"
								value={formData.name}
								placeholder="Enter name"
								onChange={handleChange}
								required
							/>
							{errors.name && <p>{errors.name}</p>}
						</div>
						<div className="input_box">
							<label>Contact Number</label>
							<input
								type="text"
								name="contact_number"
								value={formData.contact_number}
								placeholder="Enter contact number"
								onChange={handleChange}
								required
							/>
							{errors.contact_number && (
								<p>{errors.contact_number}</p>
							)}
						</div>
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
							<label>Password</label>
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
						<div className="input_box">
							<label>Confirm Password</label>
							<input
								type="password"
								name="confirmPassword"
								value={formData.confirmPassword}
								placeholder="Enter your confirm password"
								onChange={handleChange}
								required
							/>
							{errors.confirmPassword && (
								<p>{errors.confirmPassword}</p>
							)}
						</div>
						<button type="submit">SIGNUP</button>
						<p className="log_in">
							Already have an account?{" "}
							<Link to="/login">LOGIN</Link>
						</p>
					</form>
				</div>
			</div>
		</>
	);
}

export default Signup;
