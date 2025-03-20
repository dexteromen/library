import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import "./CreateLibrary.css";
import Navbar from "../../Components/Navbar/Navbar";
import { createLibrary, refreshToken } from "../../API/API";

function CreateLibrary() {
	const navigate = useNavigate();
	const [formData, setFormData] = useState({
		email: "",
		library_name: "",
	});
	const [errors, setErrors] = useState({});
	const [isSubmitting, setIsSubmitting] = useState(false);

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

		if (!formData.library_name)
			newErrors.library_name = "Library name is required";

		if (formData.library_name && formData.library_name.length < 5)
			newErrors.library_name =
				"Library name must be at least 5 characters long";

		return newErrors;
	};

	const handleSubmit = (event) => {
		event.preventDefault();
		const newErrors = validate();
		if (Object.keys(newErrors).length > 0) {
			setErrors(newErrors);
		} else {
			setErrors({});
			setIsSubmitting(true);
		}
	};

	useEffect(() => {
		if (!isSubmitting) return;

		async function fetchData() {
			try {
				const res = await createLibrary(formData.library_name);
				console.log("Library Created Successfully");
				console.log(res);
				navigate("/create-book");
			} catch (error) {
				console.log(error);
			} finally {
				setIsSubmitting(false);
			}
		}
		fetchData();

		const refresh = async () => {
			try {
				await refreshToken();
				console.log("Token refreshed successfully");
			} catch (error) {
				console.error("Failed to refresh token:", error);
			}
		};

		refresh();
	}, [isSubmitting, formData]);

	return (
		<>
			<Navbar />
			<div className="create-book-wrapper">
				<div className="create-book-form-container">
					<h1 className="create-book-title">Create Library</h1>
					<form onSubmit={handleSubmit}>
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

// import React, { useState, useEffect } from "react";
// import "./CreateLibrary.css";
// import Navbar from "../../Components/Navbar/Navbar";
// import { createLibrary } from "../../API/API";

// function CreateLibrary() {
// 	const [formData, setFormData] = useState({
// 		email: "",
// 		// libid: "",
// 		library_name: "",
// 	});
// 	const [errors, setErrors] = useState({});

// 	const handleChange = (event) => {
// 		const { name, value } = event.target;
// 		setFormData({
// 			...formData,
// 			[name]: value,
// 		});
// 	};

// 	const validate = () => {
// 		const newErrors = {};

// 		if (!formData.email) newErrors.email = "User Id is required";

// 		// if (!formData.libid) newErrors.libid = "Library Id is not provided";

// 		// if (formData.libid && formData.libid.length < 2)
// 		// 	newErrors.libid = "Library Id must be at least 2 characters long";

// 		if (!formData.library_name)
// 			newErrors.library_name = "Library name is required";

// 		if (formData.library_name && formData.library_name.length < 5)
// 			newErrors.library_name =
// 				"Library name must be atleast 5 character long";

// 		return newErrors;
// 	};

// 	const handleSubmit = (event) => {
// 		event.preventDefault();
// 		const newErrors = validate();
// 		if (Object.keys(newErrors).length > 0) {
// 			setErrors(newErrors);
// 		} else {
// 			setErrors({});
// 			// alert("Library Created Successfully!");
// 			console.log("Library Created Successfully");
// 			console.log(formData);
// 			console.log(formData.library_name);
// 			console.log(formData.email);
// 		}
// 	};

// 	useEffect(() => {
// 		async function fetchData() {
// 			try {
// 				const res = await createLibrary(formData.library_name);
// 				console.log(res);
// 			} catch (error) {
// 				console.log(error);
// 			}
// 		}
// 		fetchData();
// 	}, []);
// 	return (
// 		<>
// 			<Navbar />
// 			<div className="create-book-wrapper">
// 				<div className="create-book-form-container">
// 					<h1 className="create-book-title">Create Library</h1>
// 					<form onSubmit={handleSubmit}>
// 						{/* <div className="form-group">
// 							<label htmlFor="libid">Library Id</label>
// 							<input
// 								type="number"
// 								name="libid"
// 								id="libid"
// 								placeholder="Enter Library Id"
// 								value={formData.libid}
// 								onChange={handleChange}
// 							/>
// 							{errors.libid && (
// 								<span className="error">{errors.libid}</span>
// 							)}
// 						</div> */}

// 						<div className="form-group">
// 							<label htmlFor="email">Email</label>
// 							<input
// 								type="email"
// 								name="email"
// 								id="email"
// 								placeholder="Enter Email"
// 								value={formData.email}
// 								onChange={handleChange}
// 							/>
// 							{errors.email && (
// 								<span className="error">{errors.email}</span>
// 							)}
// 						</div>

// 						<div className="form-group">
// 							<label htmlFor="library_name">Library Name</label>
// 							<input
// 								type="text"
// 								name="library_name"
// 								id="library_name"
// 								placeholder="Enter Library Name"
// 								value={formData.library_name}
// 								onChange={handleChange}
// 							/>
// 							{errors.library_name && (
// 								<span className="error">
// 									{errors.library_name}
// 								</span>
// 							)}
// 						</div>

// 						<div className="btn-submit">
// 							<button className="button-59" type="submit">
// 								Submit
// 							</button>
// 						</div>
// 					</form>
// 				</div>
// 			</div>
// 		</>
// 	);
// }

// export default CreateLibrary;
