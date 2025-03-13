import React, { useState } from "react";
import { Link } from "react-router-dom";
import "./FormTemplate.css"

function FormTemplate({ title, fields, onSubmit, linkText, linkTo, linkValue, validate }) {
	const [formData, setFormData] = useState(
		fields.reduce((acc, field) => ({ ...acc, [field.name]: "" }), {})
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
									<p className="errors">{errors[field.name]}</p>
								)}
							</div>
						))}
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
