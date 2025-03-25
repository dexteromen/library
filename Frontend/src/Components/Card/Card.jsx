// import React from "react";
// import "./Card.css";
// import { createRequest, deleteBookById } from "../../API/API";
// import { ToastContainer, toast } from "react-toastify";
// import "react-toastify/dist/ReactToastify.css";

// function Card({
//     isbn,
//     title,
//     author,
//     publisher,
//     version,
//     lib_id,
//     total_copies,
//     available_copies,
//     user,
//     updates,
// }) {
//     const handleRequestClick = async () => {
//         if (window.confirm("Do you want to request this book?")) {
//             try {
//                 const res = await createRequest({ isbn });
//                 // alert("Request successful!");
//                 // console.log(res);
//                 toast.success("Request successfully !!");
//             } catch (error) {
//                 // alert("Request failed. Please try again.");
//                 // console.log(error.response.data);
//                 var err = error.response.data.error;
//                 toast.error(err);
//                 var message = error.response.data.message;
//                 toast.error(message);
//             }
//         }
//     };
//     const handleDeleteBook = async () => {
//         if (window.confirm("Do you want to Delete this book?")) {
//             try {
//                 // console.log(typeof isbn);
//                 const res = await deleteBookById(isbn);
//                 // alert("Request successful!");
//                 console.log(res);
//                 toast.success("Book Deleted successfully !!");
//                 setTimeout(() => {
//                     updates();
//                 }, 1000);
//                 // Reload the page
//                 // 	window.location.reload();
//             } catch (error) {
//                 // alert("Request failed. Please try again.");
//                 // console.log(error.response.data);
//                 var err = error.response.data.error;
//                 toast.error(err);
//                 var message = error.response.data.message;
//                 toast.error(message);
//             }
//         }
//     };

//     return (
//         <>
//             <ToastContainer position="top-center" />
//             <div className="card">
//                 <table className="book-details">
//                     <tbody>
//                         <tr>
//                             <td className="table-label">ISBN</td>
//                             <td>{isbn}</td>
//                         </tr>
//                         <tr>
//                             <td className="table-label">Title</td>
//                             <td>{title}</td>
//                         </tr>
//                         <tr>
//                             <td className="table-label">Author</td>
//                             <td>{author}</td>
//                         </tr>
//                         <tr>
//                             <td className="table-label">Publisher</td>
//                             <td>{publisher}</td>
//                         </tr>
//                         <tr>
//                             <td className="table-label">Version</td>
//                             <td>{version}</td>
//                         </tr>
//                         <tr>
//                             <td className="table-label">Library ID</td>
//                             <td>{lib_id}</td>
//                         </tr>
//                         <tr>
//                             <td className="table-label">Total Copies</td>
//                             <td>{total_copies}</td>
//                         </tr>
//                         <tr>
//                             <td className="table-label">Available Copies</td>
//                             <td>{available_copies}</td>
//                         </tr>
//                     </tbody>
//                 </table>
//                 {user.role === "reader" && (
//                     <button
//                         className="btn-request"
//                         onClick={handleRequestClick}
//                     >
//                         Request
//                     </button>
//                 )}
//                 {(user.role === "admin" || user.role === "owner") && (
//                     <button
//                         className="button-update-book"
//                         // onClick={handleUpdateBook}
//                     >
//                         Update Book
//                     </button>
//                 )}
//                 {(user.role === "admin" || user.role === "owner") && (
//                     <button
//                         className="button-delete-book"
//                         onClick={handleDeleteBook}
//                     >
//                         Delete Book
//                     </button>
//                 )}
//             </div>
//         </>
//     );
// }

// export default Card;

import React, { useState } from "react";
import "./Card.css";
import { createRequest, deleteBookById, updateBookById } from "../../API/API";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

function Card({
	isbn,
	title,
	author,
	publisher,
	version,
	lib_id,
	total_copies,
	available_copies,
	user,
	updates,
}) {
	const [showOverlay, setShowOverlay] = useState(false);
	const [formData, setFormData] = useState({
		isbn,
		title,
		author,
		publisher,
		version,
		lib_id,
		total_copies,
		available_copies,
	});

	const handleRequestClick = async () => {
		if (window.confirm("Do you want to request this book?")) {
			try {
				await createRequest({ isbn });
				toast.success("Request successfully submitted!");
			} catch (error) {
				const errMessage =
					error?.response?.data?.error ||
					"Request failed. Please try again.";
				toast.error(errMessage);
			}
		}
	};

	const handleDeleteBook = async () => {
		if (window.confirm("Do you want to delete this book?")) {
			try {
				await deleteBookById(isbn);
				toast.success("Book deleted successfully!");
				setTimeout(() => updates(), 1000);
			} catch (error) {
				const errMessage =
					error?.response?.data?.error ||
					"Deletion failed. Please try again.";
				toast.error(errMessage);
			}
		}
	};

	const handleUpdateBook = () => {
		setShowOverlay(true);
	};

	const handleFormChange = (e) => {
		const { name, value } = e.target;
		setFormData((prevData) => ({
			...prevData,
			[name]: value,
		}));
	};

	const validateForm = () => {
		const { title, author, total_copies, available_copies } = formData;
		if (!title.trim()) return "Title is required.";
		if (!author.trim()) return "Author is required.";
		if (isNaN(total_copies) || total_copies <= 0)
			return "Total Copies must be a positive number.";
		if (isNaN(available_copies) || available_copies < 0)
			return "Available Copies cannot be negative.";
		if (available_copies > total_copies)
			return "Available Copies cannot exceed Total Copies.";
		return null;
	};

	const handleFormSubmit = async (e) => {
		e.preventDefault();
		const validationError = validateForm();
		if (validationError) {
			toast.error(validationError);
			return;
		}
		try {
			await updateBookById(isbn, formData);
			toast.success("Book updated successfully!");
			setShowOverlay(false);
			updates();
		} catch (error) {
			const errMessage =
				error?.response?.data?.error ||
				"Update failed. Please try again.";
			toast.error(errMessage);
		}
	};

	return (
		<>
			<ToastContainer position="top-center" />
			<div className="card">
				<table className="book-details">
					<tbody>
						<tr>
							<td className="table-label">ISBN</td>
							<td>{isbn}</td>
						</tr>
						<tr>
							<td className="table-label">Title</td>
							<td>{title}</td>
						</tr>
						<tr>
							<td className="table-label">Author</td>
							<td>{author}</td>
						</tr>
						<tr>
							<td className="table-label">Publisher</td>
							<td>{publisher}</td>
						</tr>
						<tr>
							<td className="table-label">Version</td>
							<td>{version}</td>
						</tr>
						<tr>
							<td className="table-label">Library ID</td>
							<td>{lib_id}</td>
						</tr>
						<tr>
							<td className="table-label">Total Copies</td>
							<td>{total_copies}</td>
						</tr>
						<tr>
							<td className="table-label">Available Copies</td>
							<td>{available_copies}</td>
						</tr>
					</tbody>
				</table>
				{user.role === "reader" && (
					<button
						className="btn-request"
						onClick={handleRequestClick}
					>
						Request
					</button>
				)}
				{user.role === "owner" && (
					<button
						className="button-update-book"
						onClick={handleUpdateBook}
					>
						Update Book
					</button>
				)}
				{(user.role === "admin" || user.role === "owner") && (
					<button
						className="button-delete-book"
						onClick={handleDeleteBook}
					>
						Delete Book
					</button>
				)}
			</div>

			{showOverlay && (
				<div className="overlay">
					<div className="overlay-content">
						<h2>Edit Book Details</h2>
						<form onSubmit={handleFormSubmit}>
							<label>
								Title:
								<input
									type="text"
									name="title"
									value={formData.title}
									onChange={handleFormChange}
								/>
							</label>
							<label>
								Author:
								<input
									type="text"
									name="author"
									value={formData.author}
									onChange={handleFormChange}
								/>
							</label>
							<label>
								Publisher:
								<input
									type="text"
									name="publisher"
									value={formData.publisher}
									onChange={handleFormChange}
								/>
							</label>
							<label>
								Version:
								<input
									type="text"
									name="version"
									value={formData.version}
									onChange={handleFormChange}
								/>
							</label>
							<label>
								Total Copies:
								<input
									type="number"
									name="total_copies"
									value={formData.total_copies}
									onChange={handleFormChange}
								/>
							</label>
							<label>
								Available Copies:
								<input
									type="number"
									name="available_copies"
									value={formData.available_copies}
									onChange={handleFormChange}
								/>
							</label>
							<button type="submit">Save Changes</button>
							<button
								type="button"
								onClick={() => setShowOverlay(false)}
							>
								Cancel
							</button>
						</form>
					</div>
				</div>
			)}
		</>
	);
}

export default Card;
