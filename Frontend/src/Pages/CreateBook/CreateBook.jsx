import React, { useState, useEffect } from "react";
import { useNavigate, Link } from "react-router-dom";
import "./CreateBook.css";
import Navbar from "../../Components/Navbar/Navbar";
import { createBook, refreshToken } from "../../API/API";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import { MdOutlineKeyboardBackspace } from "react-icons/md";

function CreateBook() {
    const dummyBook = {
        isbn: "100-100-100-101",
        title: "English Reader 2",
        authors: "Reader 2",
        publisher: "Reader 2",
        version: "1st",
        total_copies: 2,
        available_copies: 0,
    };
    const [formData, setFormData] = useState({
        isbn: "",
        // libid: "",
        title: "",
        authors: "",
        publisher: "",
        version: "",
        total_copies: "",
        available_copies: "",
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

        // Ensure ISBN is present and properly formatted
        if (
            !formData.isbn ||
            typeof formData.isbn !== "string" ||
            formData.isbn.trim() === ""
        )
            newErrors.isbn = "ISBN is required";
        else if (formData.isbn.length < 10 || formData.isbn.length > 13)
            newErrors.isbn = "ISBN must be 10 or 13 characters long";

        // Library ID validation (Uncomment if required)
        // if (!formData.libid || typeof formData.libid !== "string" || formData.libid.trim() === "")
        // 	newErrors.libid = "Library ID is required";
        // else if (formData.libid.length < 6)
        // 	newErrors.libid = "Library ID must be at least 6 characters long";

        // Title validation
        if (
            !formData.title ||
            typeof formData.title !== "string" ||
            formData.title.trim() === ""
        )
            newErrors.title = "Title is required";

        // Author validation (assuming it's a string; can also be an array)
        if (
            !formData.authors ||
            typeof formData.authors !== "string" ||
            formData.authors.trim() === ""
        )
            newErrors.authors = "Author(s) is required";

        // Publisher validation
        if (
            !formData.publisher ||
            typeof formData.publisher !== "string" ||
            formData.publisher.trim() === ""
        )
            newErrors.publisher = "Publisher is required";

        // Version validation
        if (
            !formData.version ||
            typeof formData.version !== "string" ||
            formData.version.trim() === ""
        )
            newErrors.version = "Version is required";

        // Total copies validation
        if (
            !formData.total_copies ||
            isNaN(formData.total_copies) ||
            formData.total_copies <= 0
        )
            newErrors.total_copies = "Total copies must be a positive number";

        // Available copies validation
        if (
            !formData.available_copies ||
            isNaN(formData.available_copies) ||
            formData.available_copies < 0
        )
            newErrors.available_copies =
                "Available copies must be a non-negative number";

        // Ensure available copies do not exceed total copies
        if (
            formData.total_copies &&
            formData.available_copies &&
            formData.available_copies <= formData.total_copies
        )
            newErrors.available_copies =
                "Available copies cannot exceed total copies";

        return newErrors;
    };

    // const validate = () => {
    // 	const newErrors = {};
    // 	if (!formData.isbn) newErrors.isbn = "ISBN is required";
    // 	// if (!formData.libid) newErrors.libid = "Library Id is not provided";
    // 	// if (formData.libid && formData.libid.length < 6)
    // 	// 	newErrors.libid = "Library Id must be at least 6 characters long";
    // 	if (!formData.title) newErrors.title = "Title is required";
    // 	if (!formData.authors) newErrors.authors = "Author is required";
    // 	if (!formData.publisher) newErrors.publisher = "Publisher is required";
    // 	if (!formData.version) newErrors.version = "Version is required";
    // 	if (!formData.total_copies)
    // 		newErrors.total_copies = "Total Copies is required";
    // 	if (!formData.available_copies)
    // 		newErrors.available_copies = "Available Copies is required";
    // 	return newErrors;
    // };

    const handleSubmit = (event) => {
        event.preventDefault();
        const newErrors = validate();
        if (Object.keys(newErrors).length > 0) {
            setErrors(newErrors);
        } else {
            setErrors({});
            setIsSubmitting(true);
            console.log("Form submitted successfully");
            // console.log(formData);
        }
    };

    useEffect(() => {
        if (!isSubmitting) return;
        async function fetchData() {
            try {
                const res = await createBook(formData);
                toast.success("Book Created Successfully");
                // console.log("Book Created Successfully");
                // console.log(res);
                // navigate("/create-book");
            } catch (error) {
                console.log(error);
                // var err = error.response.data.data.Detail;
                // toast.error(err);
                // console.log(err);
            } finally {
                setIsSubmitting(false);
            }
        }
        fetchData();
    }, [isSubmitting, formData]);

    return (
        <>
            <Navbar />
            <div className="create-book-wrapper">
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
                <div className="create-book-form-container">
                    <h1 className="create-book-title">Create Book</h1>
                    <form onSubmit={handleSubmit}>
                        {/* <div className="form-group">
							<label htmlFor="libid">Library Id</label>
							<input
								type="text"
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
                            <label htmlFor="isbn">ISBN</label>
                            <input
                                type="text"
                                name="isbn"
                                id="isbn"
                                placeholder="Enter ISBN Number"
                                value={formData.isbn}
                                onChange={handleChange}
                            />
                            {errors.isbn && (
                                <span className="error">{errors.isbn}</span>
                            )}
                        </div>

                        <div className="form-group">
                            <label htmlFor="title">Title</label>
                            <input
                                type="text"
                                name="title"
                                id="title"
                                placeholder="Enter Book Title"
                                value={formData.title}
                                onChange={handleChange}
                            />
                            {errors.title && (
                                <span className="error">{errors.title}</span>
                            )}
                        </div>
                        <div className="form-group">
                            <label htmlFor="authors">Author</label>
                            <input
                                type="text"
                                name="authors"
                                id="authors"
                                placeholder="Enter Book Author"
                                value={formData.authors}
                                onChange={handleChange}
                            />
                            {errors.authors && (
                                <span className="error">{errors.authors}</span>
                            )}
                        </div>
                        <div className="form-group">
                            <label htmlFor="publisher">Publisher</label>
                            <input
                                type="text"
                                name="publisher"
                                id="publisher"
                                placeholder="Enter Book Publisher"
                                value={formData.publisher}
                                onChange={handleChange}
                            />
                            {errors.publisher && (
                                <span className="error">
                                    {errors.publisher}
                                </span>
                            )}
                        </div>
                        <div className="form-group">
                            <label htmlFor="version">Version</label>
                            <input
                                type="text"
                                name="version"
                                id="version"
                                placeholder="Enter Book Version"
                                value={formData.version}
                                onChange={handleChange}
                            />
                            {errors.version && (
                                <span className="error">{errors.version}</span>
                            )}
                        </div>
                        <div className="form-group">
                            <label htmlFor="total_copies">Total Copies</label>
                            <input
                                type="text"
                                name="total_copies"
                                id="total_copies"
                                placeholder="Enter Total Copies"
                                value={formData.total_copies}
                                onChange={handleChange}
                            />
                            {errors.total_copies && (
                                <span className="error">
                                    {errors.total_copies}
                                </span>
                            )}
                        </div>
                        <div className="form-group">
                            <label htmlFor="available_copies">
                                Available Copies
                            </label>
                            <input
                                type="text"
                                name="available_copies"
                                id="available_copies"
                                placeholder="Enter Available Copies"
                                value={formData.available_copies}
                                onChange={handleChange}
                            />
                            {errors.available_copies && (
                                <span className="error">
                                    {errors.available_copies}
                                </span>
                            )}
                        </div>
                        {/* <button className="book-submit-btn" type="submit">
							Submit
						</button> */}
                        <div className="btn-submit">
                            <button className="button-59" type="submit">
                                Submit
                            </button>
                        </div>
                    </form>
                    <div className="child-back-btn">
                        <Link to="/home">
                            <MdOutlineKeyboardBackspace size={45} />
                        </Link>
                    </div>
                    {/* <div className="multi-options">
						<Link to="/home">
							<button>Home</button>
						</Link>
					</div> */}
                </div>
            </div>
        </>
    );
}

export default CreateBook;
