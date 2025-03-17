import React, { useState } from "react";
import "./CreateLibrary.css";
import Navbar from "../../Components/Navbar/Navbar";

function CreateLibrary() {
	const [formData, setFormData] = useState({
		email: "",
		// libid: "",
		library_name: "",
	});
	const [errors, setErrors] = useState({});

	const handleChange = (event) => {
		const { name, value } = event.target;
		setFormData({
			...formData,
			[name]: value,
		});
	};

	const validate = () => {
		const newErrors = {};

		if (!formData.email) newErrors.email = "User Id is required";

		// if (!formData.libid) newErrors.libid = "Library Id is not provided";

		// if (formData.libid && formData.libid.length < 2)
		// 	newErrors.libid = "Library Id must be at least 2 characters long";

		if (!formData.library_name)
			newErrors.library_name = "Library name is required";

		if (formData.library_name && formData.library_name.length < 5)
			newErrors.library_name =
				"Library name must be atleast 5 character long";

		return newErrors;
	};

	const handleSubmit = (event) => {
		event.preventDefault();
		const newErrors = validate();
		if (Object.keys(newErrors).length > 0) {
			setErrors(newErrors);
		} else {
			setErrors({});
			alert("Library Created Successfully!");
			console.log("Form submitted successfully");
			console.log(formData);
		}
	};
	return (
		<>
			<Navbar />
			<div className="create-book-wrapper">
				<div className="create-book-form-container">
					<h1 className="create-book-title">Create Library</h1>
					<form onSubmit={handleSubmit}>
						{/* <div className="form-group">
							<label htmlFor="libid">Library Id</label>
							<input
								type="number"
								name="libid"
								id="libid"
								placeholder="Enter Library Id"
								value={formData.libid}
								onChange={handleChange}
							/>
							{errors.libid && (
								<span className="error">{errors.libid}</span>
							)}
						</div> */}

						<div className="form-group">
							<label htmlFor="email">Email</label>
							<input
								type="email"
								name="email"
								id="email"
								placeholder="Enter Email"
								value={formData.email}
								onChange={handleChange}
							/>
							{errors.email && (
								<span className="error">{errors.email}</span>
							)}
						</div>

						<div className="form-group">
							<label htmlFor="library_name">Library Name</label>
							<input
								type="text"
								name="library_name"
								id="library_name"
								placeholder="Enter Library Name"
								value={formData.library_name}
								onChange={handleChange}
							/>
							{errors.library_name && (
								<span className="error">
									{errors.library_name}
								</span>
							)}
						</div>

						<div className="btn-submit">
							<button className="button-59" type="submit">
								Submit
							</button>
						</div>
					</form>
				</div>
			</div>
		</>
	);
}

export default CreateLibrary;
