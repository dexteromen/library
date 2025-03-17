import React from "react";
import FormTemplate from "../../Components/FormTemplate/FormTemplate";
import Navbar from "../../Components/Navbar/Navbar";

function Signup() {
	const SignupFields = [
		{
			name: "name",
			label: "Name",
			type: "text",
			placeholder: "Enter name",
			required: true,
		},
		{
			name: "contact_number",
			label: "Contact Number",
			type: "text",
			placeholder: "Enter contact number",
			required: true,
		},
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
		{
			name: "confirmPassword",
			label: "Confirm Password",
			type: "password",
			placeholder: "Enter your confirm password",
			required: true,
		},
	];

	const validate = (formData) => {
		let tempErrors = {};
		if (!formData.name) tempErrors.name = "Name is required";
		if (formData.contact_number.length !== 10)
			tempErrors.contact_number = "Contact number must be 10 digits";
		if (!formData.email) tempErrors.email = "Email is required";
		if (!/\S+@\S+\.\S+/.test(formData.email))
			tempErrors.email = "Email is invalid";
		if (!formData.password) tempErrors.password = "Password is required";
		if (formData.password.length < 6)
			tempErrors.password = "Password must be at least 6 characters";
		if (formData.password !== formData.confirmPassword)
			tempErrors.confirmPassword = "Passwords do not match";
		return tempErrors;
	};

	const handleSubmit = (formData) => {
		console.log("Form submitted successfully");
		console.log(formData);
	};
	return (
		<>
			<div className="container-signup">
				<Navbar />
				<FormTemplate
					title="SIGNUP"
					fields={SignupFields}
					onSubmit={handleSubmit}
					validate={validate}
					linkText="Already have an account?"
					linkTo="/login"
					linkValue="LOGIN"
				/>
			</div>
		</>
	);
}

export default Signup;
