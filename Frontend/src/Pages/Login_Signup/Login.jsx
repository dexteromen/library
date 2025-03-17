import React from "react";
import FormTemplate from "../../Components/FormTemplate/FormTemplate";
import Navbar from "../../Components/Navbar/Navbar";

function Login() {
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

	const handleSubmit = (formData) => {
		console.log("Form submitted successfully");
		console.log(formData);
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
				/>
			</div>
		</>
	);
}

export default Login;
