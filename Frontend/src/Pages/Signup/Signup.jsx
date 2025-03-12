import React, { useState } from "react";
import Navbar from "../../Components/Navbar/Navbar";

function Signup() {
	const [formData, setFormData] = useState({
		username: "",
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
		if (!formData.username) tempErrors.username = "Username is required";
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
			<div className="wrapper-signup">
				<form className="form" onSubmit={handleSubmit}>
					<div className="field-div">
						<label>Username</label>
						<input
							type="text"
							name="username"
							value={formData.username}
							onChange={handleChange}
						/>
						{errors.username && <p>{errors.username}</p>}
					</div>
					<div className="field-div">
						<label>Email</label>
						<input
							type="email"
							name="email"
							value={formData.email}
							onChange={handleChange}
						/>
						{errors.email && <p>{errors.email}</p>}
					</div>
					<div className="field-div">
						<label>Password</label>
						<input
							type="password"
							name="password"
							value={formData.password}
							onChange={handleChange}
						/>
						{errors.password && <p>{errors.password}</p>}
					</div>
					<div className="field-div">
						<label>Confirm Password</label>
						<input
							type="password"
							name="confirmPassword"
							value={formData.confirmPassword}
							onChange={handleChange}
						/>
						{errors.confirmPassword && (
							<p>{errors.confirmPassword}</p>
						)}
					</div>
					<button type="submit">Signup</button>
				</form>
			</div>
		</>
	);
}

export default Signup;
