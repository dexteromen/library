import React from "react";
import "./Card.css";
import { createRequest } from "../../API/API";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

function Card({ isbn, title, author, publisher, version }) {
	const handleRequestClick = async () => {
		if (window.confirm("Do you want to request this book?")) {
			try {
				const res = await createRequest({ isbn });
				// alert("Request successful!");
				// console.log(res);
				toast.success("Request successfully !!");
			} catch (error) {
				// alert("Request failed. Please try again.");
				// console.log(error.response.data);
				var err = error.response.data.error;
				toast.error(err);
				var message = error.response.data.message;
				toast.error(message);
			}
		}
	};

	return (
		<>
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
					</tbody>
				</table>
				<button className="btn-request" onClick={handleRequestClick}>
					Request
				</button>
			</div>
		</>
	);
}

export default Card;
