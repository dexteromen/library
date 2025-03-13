import React, { useState, useEffect } from "react";
import "./CreateBook.css";
import Navbar from "../../Components/Navbar/Navbar";

function CreateBook() {
	const [formData, setFormData] = useState({});
	const [errors, setErrors] = useState({});

	const fields = [
		{
			name: "isbn",
			label: "ISBN",
			type: "text",
			placeholder: "Enter ISBN Number",
			required: true,
		},
		{
			name: "libid",
			label: "Library Id",
			type: "text",
			placeholder: "Enter Library Id",
			required: true,
		},
		{
			name: "title",
			label: "Title",
			type: "text",
			placeholder: "Enter Book Title",
			required: true,
		},
		{
			name: "authors",
			label: "Author",
			type: "text",
			placeholder: "Enter Book Author",
			required: true,
		},
		{
			name: "publisher",
			label: "Publisher",
			type: "text",
			placeholder: "Enter Book Publisher",
			required: true,
		},
		{
			name: "version",
			label: "Version",
			type: "text",
			placeholder: "Enter Book Version",
			required: true,
		},
		{
			name: "total_copies",
			label: "Total Copies",
			type: "text",
			placeholder: "Enter Total Copies",
			required: true,
		},
		{
			name: "available_copies",
			label: "Available Copies",
			type: "text",
			placeholder: "Enter Available Copies",
			required: true,
		},
	];

	const handleChange = (event) => {
		const { name, value } = event.target;
		setFormData({
			...formData,
			[name]: value,
		});
	};

	const validate = () => {
		const newErrors = {};
		fields.forEach((field) => {
			if (field.required && !formData[field.name]) {
				newErrors[field.name] = `${field.label} is required`;
			}
			if (
				field.name === "libid" &&
				formData[field.name] &&
				formData[field.name].length < 6
			) {
				newErrors[field.name] =
					"Library Id must be at least 6 characters long";
			}
		});
		return newErrors;
	};

	const handleSubmit = (event) => {
		event.preventDefault();
		const newErrors = validate();
		if (Object.keys(newErrors).length > 0) {
			setErrors(newErrors);
		} else {
			setErrors({});
			console.log("Form submitted successfully");
			console.log(formData);
		}
	};

	return (
		<>
			<Navbar />
			<div className="form-container">
				<form onSubmit={handleSubmit}>
					{fields.map((field) => (
						<div key={field.name} className="form-group">
							<label htmlFor={field.name}>{field.label}</label>
							<input
								type={field.type}
								name={field.name}
								id={field.name}
								placeholder={field.placeholder}
								required={field.required}
								onChange={handleChange}
							/>
							{errors[field.name] && (
								<span className="error">
									{errors[field.name]}
								</span>
							)}
						</div>
					))}
					<button className="book-submit-btn" type="submit">
						Submit
					</button>
				</form>
			</div>
		</>
	);
}

export default CreateBook;
