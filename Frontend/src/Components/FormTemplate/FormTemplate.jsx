import React, { useState } from "react";
import { Link } from "react-router-dom";
import "./FormTemplate.css";

function FormTemplate({
	title,
	fields,
	onSubmit,
	linkText,
	linkTo,
	linkValue,
	validate,
	roleSelector,
}) {
	const [formData, setFormData] = useState(
		fields.reduce(
			(acc, field) => ({
				...acc,
				[field.name]: "",
			}),
			{ role: "" }
		)
	);
	const [errors, setErrors] = useState({});

	const handleChange = (e) => {
		const { name, value } = e.target;
		setFormData({
			...formData,
			[name]: value,
		});
	};

	const handleSubmit = (e) => {
		e.preventDefault();
		const tempErrors = validate(formData);
		setErrors(tempErrors);
		if (Object.keys(tempErrors).length === 0) {
			onSubmit(formData);
		}
	};

	return (
		<>
			<div className="templete-form-wrapper">
				<div className="templete-form-parent">
					<form className="templete-form" onSubmit={handleSubmit}>
						<h1>{title}</h1>
						{fields.map((field) => (
							<div className="input_box" key={field.name}>
								<label>{field.label}</label>
								<input
									type={field.type}
									name={field.name}
									value={formData[field.name]}
									placeholder={field.placeholder}
									onChange={handleChange}
									required={field.required}
								/>
								{errors[field.name] && (
									<p className="errors">
										{errors[field.name]}
									</p>
								)}
							</div>
						))}
						{roleSelector === true && (
							<div className="role-selector">
								<select
									name="role"
									value={formData.role || ""}
									onChange={handleChange}
									required
								>
									<option value="" disabled>
										Select a role
									</option>
									<option value="reader">Reader</option>
									<option value="admin">Admin</option>
									{/* <option value="owner">Owner</option> */}
								</select>
							</div>
						)}
						<button type="submit">{title}</button>
						<p className="link-to">
							{linkText} <Link to={linkTo}>{linkValue}</Link>
						</p>
					</form>
				</div>
			</div>
		</>
	);
}

export default FormTemplate;
