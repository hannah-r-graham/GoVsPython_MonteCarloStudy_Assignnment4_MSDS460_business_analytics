# GoVsPython_MonteCarloStudy_Assignnment4_MSDS460_business_analytics



For this assignment, we will use simulated data and run a performance benchmark. You are welcome to use Go, Python, and/or R for this assignment. In fact, you are encouraged to work on a case study with accompanying data using each of the three languages of data science: Go, Python, and R.

Your team is welcome to define a performance benchmark of its own (perhaps something that flows from another data science course or from your business experience). Just ensure that the research study your team proposes employs Monte Carlo simulation for data generation (with at least two levels or types of data) and compares at least two database systems, programming languages, or modeling methods.

The data sets that your workgroup uses for this assignment can be computer-generated using random number generators from your selected language(s) or data from an open-source data source. This short video cites 10 Free Dataset Resources for Your Next ProjectLinks to an external site..

The core of your research should involve a comparison between at least two database systems, programming languages, or modeling tools. Set up a benchmark study comparing performance across the systems or tools.  We use the term "performance" here to refer to either computer performance (processing time and memory utilization) or predictive performance (in classification or regression models).

Note that batch/command-line execution is best for performance benchmarks. 

In addition to summarizing performance results, your paper should answer these questions:

How long did it take to implement the benchmark?
How long did it take to bring data into the systems? 
How long did it take to get data out of the systems?
In addition to the benchmark, consider the extent to which companies utilize these systems or tools. What are their market shares? What is the status of documentation and training materials regarding the systems or tools you have selected? 
How difficult is it for a data scientist like yourself to work with these systems?
What is the hypothetical management problem that your benchmark study addresses?
What would it cost to implement the database systems, programs, or modeling methods being compared?
Data scientists are often called on to compare modeling methods for regression and classification problems. Hothorn et al. (2005) described benchmark methods as they are used in data science and statistics. This assignment calls for an information systems benchmark. We are concerned with computer resources required to install and execute database systems and tools or modeling methods and tools.

Many student teams may want to explore modern statistical methods as described in Gelman and Vehtari. (2022). Modern methods are typically computer-intensive methods, starting with jackknife and bootstrap resampling methods and extending through Bayesian hierarchical modeling (with Markov Chain Monte Carlo estimation of posterior distributions). Multifold cross-validation is a common technology for validating models. These fall within the domain of Monte Carlo or statistical simulation methods. Any of these methods may be explored as part of this assignment.

Many Monte Carlo benchmark studies fall within the general area of computer systems performance analysis. See References at the bottom of this write-up for additional readings in this area. 

Database systems are another good source of questions for data scientists and data engineers. Your workgroup could select one or more database systems and/or tools, perhaps comparing newer flexible or enhanced tools with a PostgreSQL/Python/psycopg2 baseline.
Deliverables (150 points total, 30 points for each part) (One submission per workgroup)

Remember to include all input data, programs/code, results/listings, and documentation in your GitHub repository.

One member of the workgroup should be designated as the workgroup leader, and this person should be posting the deliverables. Only one posting is needed for each workgroup. Members of the workgroup will share a common grade on this assignment with the understanding that all workgroup members contribute to the work.

Paper (pdf file). The paper/write-up should be submitted as a pdf file (double-spaced, 4 pages of text max, with accompanying tables and figures). Think of the paper as a complete written research report with six sections as described below. If you like, provide a paragraph or two on methods and a paragraph or two about results for each part of this assignment. 

Program code (text link to GitHub repository). Key information from the paper should also be included in the README.md markdown file of a public GitHub repository established by the student who is submitting on behalf of the workgroup.  The GitHub repository should include text files for the program code (.py extension for Python, .R for R, and/or .go for Go), and program output (.txt extension). Tables (.txt or .csv extension) and figures (.jpg or .png extension) should be included in the GitHub repository. Include the web address text (URL) for the GitHub repository in the comments form of the assignment posting.  You should be providing a link to a separate repository that can be cloned. It should end with the .git extension.

Written Report Components

Organize the paper as a research report with each section answering basic questions as identified in bold type below:

Abstract. An executive summary of the research. Describe the work that your team has completed in setting up this benchmark study of database systems.
Introduction. Why are you conducting this research?  
Literature review. Who else has conducted research like this? 
Methods. How are you conducting the research? Describe your methods for generating simulated data and running the Monte Carlo study.
Results. What did you learn from the research?
Conclusions. So, what does it all mean?  To what extent is your workgroup/firm comfortable in using Monte Carlo simulation in selecting database systems and/or tools?  What are the limitations of your benchmark study design? If you had to conduct the study again, what would you do differently?
When drawing conclusions, make sure that you relate your study results back to the data engineering, data science, or management problem you have identified. 

See the Writing Research Reports page under the MSDS Resources Module. 

Organization and File Formats of Deliverables

Deliverables

The paper summarizing your work should consist of no more than four pages of text, double-spaced. Figures and tables may go in an appendix and do not count as part of the four pages. The paper should be provided as an Adobe Acrobat pdf file. (MS Word files are NOT acceptable.) This paper comprises all sections of a written research report. It should summarize performance results and answer the questions listed at the beginning of this assignment write-up.  Note that managers are unlikely to dig into programs and output files from programs. The results portion of the written report needs to present actual results. Include tables and figures to help the reader.
Program files should be provided as plain text. Python, R, and Go are the recommended programming languages. These may be supplemented by SQL plain text files showing database queries, which are executed within the Python, R, and/or Go programs. Plain text files are needed so your instructor can execute your programs to check your work.
Include output from programs, such as console listings/logs, text files, and graphics output (pdf, png, or jpg) for visualizations. 
It will be most convenient to run the benchmark on your personal computer (localhost server).  
If you use Jupyter notebook to work on this assignment, ensure that, in addition to your ipynb script, you include a plain text .py file of your program and an html file showing the results from executing your script. Ensure that all code and visualizations are fully visible in the files that you submit.
List of file names and descriptions of files in the in a plain text file within the GitHub repository.
