import React from "react";
import FormTemplate from "../../Components/FormTemplate/FormTemplate";
import Navbar from "../../Components/Navbar/Navbar";
import { useNavigate } from "react-router-dom";
import { signIn } from "../../API/API";

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
		try {
			const creadentials = {
				email: formData.email,
				password: formData.password,
			};
			const res = await signIn(creadentials);

			// console.log(res.data.data);
			const { token,user_id } = res.data.data;
			localStorage.setItem("token", token);
			localStorage.setItem("user_id", user_id);
			// localStorage.setItem("email", creadentials.email);
			navigate("/");
			console.log("User logged-in successfully !!");
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
