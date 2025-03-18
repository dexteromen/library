import React from "react";
import FormTemplate from "../../Components/FormTemplate/FormTemplate";
import Navbar from "../../Components/Navbar/Navbar";
import { useNavigate } from "react-router-dom";
import axios from "axios";

function Login() {
	const navigate = useNavigate();

	const LoginFields = [
		{
			name: "email",
			label: "Email",
			type: "email",
			placeholder: "Enter email address",
			required: true,
		},
		{
			name: "password",
			label: "Password",
			type: "password",
			placeholder: "Enter your password",
			required: true,
		},
	];

	const validate = (formData) => {
		let tempErrors = {};
		if (!formData.email) tempErrors.email = "Email is required";
		if (!/\S+@\S+\.\S+/.test(formData.email))
			tempErrors.email = "Email is invalid";
		if (!formData.password) tempErrors.password = "Password is required";
		if (formData.password.length < 6)
			tempErrors.password = "Password must be at least 6 characters";
		return tempErrors;
	};

	const handleSubmit = async (formData) => {
		// console.log(formData);
		const URL = "http://localhost:8080/signin";
		try {
			const res = await axios.post(URL, {
				email: formData.email,
				password: formData.password,
			});
			const { token, expiry_time, user_id } = res.data.data;
			localStorage.setItem("token", token);
			localStorage.setItem("user_id", user_id);
			// localStorage.setItem("expiry_time", expiry_time);
			navigate("/");
			console.log("User logged-in successfully !!");
			// console.log(res.data);
		} catch (error) {
			console.log(error);
		}
	};

	return (
		<>
			<div className="container-login">
				<Navbar />
				<FormTemplate
					title="LOGIN"
					fields={LoginFields}
					onSubmit={handleSubmit}
					validate={validate}
					linkText="Don't have an account?"
					linkTo="/signup"
					linkValue="SIGNUP"
					roleSelector={false}
				/>
			</div>
		</>
	);
}

export default Login;
