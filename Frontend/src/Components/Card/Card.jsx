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
                const res = await createRequest({ isbn });
                toast.success("Request successfully !!");
            } catch (error) {
                var err = error.response.data.error;
                toast.error(err);
                var message = error.response.data.message;
                toast.error(message);
            }
        }
    };

    const handleDeleteBook = async () => {
        if (window.confirm("Do you want to Delete this book?")) {
            try {
                const res = await deleteBookById(isbn);
                console.log(res);
                toast.success("Book Deleted successfully !!");
                setTimeout(() => {
                    updates();
                }, 1000);
            } catch (error) {
                var err = error.response.data.error;
                toast.error(err);
                var message = error.response.data.message;
                toast.error(message);
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

    const handleFormSubmit = async (e) => {
        e.preventDefault();
        try {
            const res = await updateBookById(isbn, formData);
            console.log(res);
            toast.success("Book updated successfully!");
            setShowOverlay(false);
            updates();
        } catch (error) {
            var err = error.response.data.error;
            toast.error(err);
            var message = error.response.data.message;
            toast.error(message);
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
