import React from "react";
import FormTemplate from "../../Components/FormTemplate/FormTemplate";
import Navbar from "../../Components/Navbar/Navbar";
import { useNavigate } from "react-router-dom";
import { signUp } from "../../API/API";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

function Signup() {
	const navigate = useNavigate();

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

		// Validate Name: Only letters and spaces
		if (!formData.name) {
			tempErrors.name = "Name is required";
		} else if (!/^[A-Za-z\s]+$/.test(formData.name)) {
			tempErrors.name = "Only letters and spaces";
		}

		// Validate Contact Number: Exactly 10 digits
		if (!formData.contact_number) {
			tempErrors.contact_number = "Contact number is required";
		} else if (!/^\d{10}$/.test(formData.contact_number)) {
			tempErrors.contact_number =
				"Contact number must be exactly 10 digits";
		}

		// // Validate Email: Format
		// if (!formData.email) {
		// 	tempErrors.email = "Email is required";
		// } else if (!/\S+@\S+\.\S+/.test(formData.email)) {
		// 	tempErrors.email = "Email is invalid";
		// }

		// Validate Gmail Email: Format
		if (!formData.email) {
			tempErrors.email = "Email is required";
		} else if (!/^[a-zA-Z0-9._%+-]+@gmail\.com$/.test(formData.email)) {
			tempErrors.email = "Enter Valid Gmail Address";
		}

		// Validate Password: 8 characters, uppercase, number, special character
		if (!formData.password) {
			tempErrors.password = "Password is required";
		} else {
			// Check for minimum length
			if (formData.password.length < 8) {
				tempErrors.password =
					"Password must be at least 8 characters long";
			}

			// Check for at least one uppercase letter
			else if (!/[A-Z]/.test(formData.password)) {
				tempErrors.password =
					"Password must include 1 uppercase letter";
			}

			// Check for at least one number
			else if (!/\d/.test(formData.password)) {
				tempErrors.password = "Password must include 1 number";
			}

			// Check for at least one special character
			else if (!/[@$!%*?&]/.test(formData.password)) {
				tempErrors.password =
					"Password must include 1 special character";
			}
		}

		// Validate Confirm Password: Matches Password
		if (formData.password !== formData.confirmPassword) {
			tempErrors.confirmPassword = "Passwords do not match";
		}

		return tempErrors;
	};

	// const validate = (formData) => {
	// 	let tempErrors = {};
	// 	if (!formData.name) tempErrors.name = "Name is required";

	// 	if (formData.contact_number.length !== 10)
	// 		tempErrors.contact_number = "Contact number must be 10 digits";

	// 	if (!formData.email) tempErrors.email = "Email is required";

	// 	if (!/\S+@\S+\.\S+/.test(formData.email))
	// 		tempErrors.email = "Email is invalid";

	// 	if (!formData.password) tempErrors.password = "Password is required";

	// 	if (formData.password.length < 6)
	// 		tempErrors.password = "Password must be at least 6 characters";

	// 	if (formData.password !== formData.confirmPassword)
	// 		tempErrors.confirmPassword = "Passwords do not match";

	// 	return tempErrors;
	// };

	const handleSubmit = async (formData) => {
		console.log("Form submitted successfully");
		// console.log(formData);
		try {
			const userData = {
				name: formData.name,
				email: formData.email,
				contact_number: formData.contact_number,
				password: formData.password,
				role: formData.role,
			};
			const res = await signUp(userData);
			console.log(res);
			// console.log("User created successfully !!");
			toast.success("User created successfully.");
			setTimeout(() => {
				navigate("/login");
			}, 3000);
		} catch (error) {
			var err = error.response.data.error;
			toast.error(err);

			if (
				error.response.data.message ===
				"Cannot create more than one admin"
			) {
				// console.log("cant");
				toast.error("Can Not Create One More Admin.");
			}
			// console.log(error);
		}
	};
	return (
		<>
			<div className="container-signup">
				<Navbar />
				<ToastContainer
					position="top-center"
					autoClose={2000}
					hideProgressBar={false}
					newestOnTop={false}
					closeOnClick
					rtl={false}
					pauseOnFocusLoss
					draggable
					pauseOnHover
				/>
				<FormTemplate
					title="SIGNUP"
					fields={SignupFields}
					onSubmit={handleSubmit}
					validate={validate}
					linkText="Already have an account?"
					linkTo="/login"
					linkValue="LOGIN"
					roleSelector={true}
				/>
			</div>
		</>
	);
}

export default Signup;
