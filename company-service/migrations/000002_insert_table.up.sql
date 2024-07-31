-- Insert mock data into the companies table
INSERT INTO companies (id, name, address, industry, website)
VALUES
    ('123e4567-e89b-12d3-a456-426614174000', 'TechCorp', '123 Tech Street', 'Technology', 'www.techcorp.com'),
    ('223e4567-e89b-12d3-a456-426614174001', 'HealthCorp', '456 Health Avenue', 'Healthcare', 'www.healthcorp.com');

-- Insert mock data into the departments table
INSERT INTO departments (id, name, company_id)
VALUES
    ('323e4567-e89b-12d3-a456-426614174002', 'Engineering', '123e4567-e89b-12d3-a456-426614174000'),
    ('423e4567-e89b-12d3-a456-426614174003', 'Sales', '123e4567-e89b-12d3-a456-426614174000'),
    ('523e4567-e89b-12d3-a456-426614174004', 'Research', '223e4567-e89b-12d3-a456-426614174001');

-- Insert mock data into the positions table
INSERT INTO positions (id, title, description, department_id)
VALUES
    ('623e4567-e89b-12d3-a456-426614174005', 'Software Engineer', 'Develops software applications', '323e4567-e89b-12d3-a456-426614174002'),
    ('723e4567-e89b-12d3-a456-426614174006', 'Sales Manager', 'Manages sales team and strategy', '423e4567-e89b-12d3-a456-426614174003'),
    ('823e4567-e89b-12d3-a456-426614174007', 'Research Scientist', 'Conducts scientific research', '523e4567-e89b-12d3-a456-426614174004');

-- Insert mock data into the resumes table
INSERT INTO resumes (id, user_id, education, summary)
VALUES
    ('923e4567-e89b-12d3-a456-426614174008', '023e4567-e89b-12d3-a456-426614174009', 'Bachelor of Science in Computer Science', 'Experienced in software development'),
    ('a23e4567-e89b-12d3-a456-426614174010', '123e4567-e89b-12d3-a456-426614174011', 'Master of Business Administration', 'Skilled in sales and management');
