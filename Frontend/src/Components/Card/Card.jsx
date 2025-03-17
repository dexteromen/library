import React from "react";
import "./Card.css";

function Card({ bookImage = "Book Image", desc = "Description" }) {
	return (
		<>
			<div className="card">
				{/* <div className="book-image">{bookImage}</div>
				<div className="book-desc">{desc}</div> */}
				<table>
					<tbody>
						<tr>
							<td className="table-lable">ISBN</td>
							<td>ISBN of Book</td>
						</tr>
						<tr>
							<td className="table-lable">Title</td>
							<td>Title of Book</td>
						</tr>
						<tr>
							<td className="table-lable">Author</td>
							<td>Author of Book</td>
						</tr>
						<tr>
							<td className="table-lable">Publisher</td>
							<td>Publisher of Book</td>
						</tr>
						<tr>
							<td className="table-lable">Version</td>
							<td>Version of Book</td>
						</tr>
					</tbody>
				</table>
				{/* <button className="btn-view-book">View Book</button> */}
				<button className="btn-request">request</button>
			</div>
		</>
	);
}

export default Card;
