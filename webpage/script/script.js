        // Function to fetch containers from the backend
        async function fetchContainers() {
            try {
                const response = await fetch('http://localhost:8080/v1/api/retreive/containers/all', {
                    method: 'GET',
                    headers: {
                        
                        'Content-Type': 'application/json',
                        'FreshView-API-Key':CONFIG.FRESHVIEW_API_KEY
                    },

                });
                const data = await response.json();
                console.log(data)
                return data.containerNames; 
            } catch (error) {
                console.error('Error fetching containers:', error);
            }
        }

        // Function to fetch duties for a specific container
        async function fetchDuties(containerName) {
            try {
                const response = await fetch(`http://localhost:8080/v1/api/retreive/container/${containerName}`, {
                    method: 'GET',
                    headers: {
                        'Content-Type': 'application/json',
                        'FreshView-API-Key': CONFIG.FRESHVIEW_API_KEY
                    }
                });
                const data = await response.json();
                console.log(data.container.Duties);
                return data.container.Duties; // Adjust based on your actual JSON structure
            } catch (error) {
                console.error(`Error fetching duties for container ${containerName}:`, error);
            }
        }

        // Function to create and append a dutie container
        function createDutieContainer(container) {
            const dutieContainerDiv = document.createElement('div');
            dutieContainerDiv.className = 'dutieContainer';
            dutieContainerDiv.setAttribute('data-name', container); // Tag with unique name
            dutieContainerDiv.textContent = container;

            // const dutieListDiv = document.createElement('div');
            // dutieListDiv.className = 'dutieList';

            // container.duties.forEach(dutie => {
            //     const dutieDiv = document.createElement('div');
            //     dutieDiv.className = 'dutie';
            //     dutieDiv.textContent = dutie;
            //     dutieListDiv.appendChild(dutieDiv);
            // });

            // dutieContainerDiv.appendChild(dutieListDiv);
            document.querySelector('.dutiesContainersList').appendChild(dutieContainerDiv);
            
            // Add click event listener to fetch and display duties
            dutieContainerDiv.addEventListener('click', async () => {
                document.querySelector('.main').classList.add('shrink');
                const duties = await fetchDuties(container);
                displayDuties(duties,container);
                });
            }

        // Function to display duties
        function displayDuties(duties,container) {
            const dutiesDisplay = document.querySelector('.dutiesDisplay');
            const dutiesList = document.getElementById('dutiesList');
            dutiesList.innerHTML = ''; // Clear previous duties

            duties.forEach(duty => {
                const dutyDiv = document.createElement('div');
                dutyDiv.className = 'duty';
                dutyDiv.setAttribute('container-name',container)

                const taskNameDiv = document.createElement('div');
                taskNameDiv.className = 'taskName';
                taskNameDiv.textContent = `${duty.TaskName}`;

                const categoryDiv = document.createElement('div');
                categoryDiv.className = 'category';
                categoryDiv.textContent = `Category: ${duty.Category}`;

                const createdAtDiv = document.createElement('div');
                console.log("czas :", duty.CreatedAt)
                createdAtDiv.className = 'createdAt';
                var timeDate = duty.CreatedAt.replace("T"," ")
                console.log(duty.CreatedAt)
                createdAtDiv.textContent = `Created At: ${new Date(timeDate).toLocaleDateString()} ${new Date(timeDate).toLocaleTimeString()}`;

                dutyDiv.appendChild(taskNameDiv);
                dutyDiv.appendChild(categoryDiv);
                dutyDiv.appendChild(createdAtDiv);

                dutiesList.appendChild(dutyDiv);
            });

            dutiesList.appendChild(createDutieTaskDiv(container))

            dutiesDisplay.classList.add('active');
        }

        function createDutieTaskDiv(container) {
            // Create the main div
            const dutieTaskDiv = document.createElement('div');
            dutieTaskDiv.className = 'duty';
            dutieTaskDiv.classList.add('toAddContainer')
            dutieTaskDiv.setAttribute('container-name',container)
        
            // Create the Task Name input field
            const taskNameInput = document.createElement('input');
            taskNameInput.type = 'text';
            taskNameInput.id = 'taskNameInput';
            taskNameInput.placeholder = 'Task Name';
        
            // Create the Category input field
            const categoryInput = document.createElement('input');
            categoryInput.type = 'text';
            categoryInput.id = 'categoryInput';
            categoryInput.placeholder = 'Category';
        
            // Create the Add Task button
            const addTaskButton = document.createElement('button');
            addTaskButton.id = 'addTaskButton';
            addTaskButton.textContent = 'Add Task';
        
            // Append the input fields and button to the main div
            addTaskButton.addEventListener('click',addTask)
            dutieTaskDiv.appendChild(taskNameInput);
            dutieTaskDiv.appendChild(categoryInput);
            dutieTaskDiv.appendChild(addTaskButton);

        
            return dutieTaskDiv
        }

        async function addTask() {
            const taskName = document.getElementById('taskNameInput').value;
            const category = document.getElementById('categoryInput').value;
            const container = document.getElementsByClassName('toAddContainer')[0].getAttribute('container-name')
            console.log("container is ",container)
        
            if (!taskName || !category) {
                alert('Please fill in both fields');
                return;
            }
        
            try {
                const response = await fetch(`http://localhost:8080/v1/api/add/task/${container}`, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                        'FreshView-API-Key': CONFIG.FRESHVIEW_API_KEY
                    },
                    body: JSON.stringify({
                        taskName: taskName,
                        taskCategory: category
                    })
                });
                console.log(response)
        
                if (response.ok) {
                    alert('Task added successfully');
                    document.getElementById('taskNameInput').value = '';
                    document.getElementById('categoryInput').value = '';
                } else {
                    const errorData = await response.json();
                    console.error('Error adding task:', errorData);
                    alert('Error adding task');
                }
            } catch (error) {
                console.error('Error adding task:', error);
                alert('Error adding task');
            }
            // displayDuties(fetchDuties(container))

        }


        // Function to initialize the process
        async function initialize() {
            const containers = await fetchContainers();
            containers.forEach(container => {
                createDutieContainer(container);
            });
        }

        // Initialize on page load
        window.onload = initialize;

        // Example function to do operations based on container names
        function doSomethingWithContainer(name) {
            const container = document.querySelector(`.dutieContainer[data-name='${name}']`);
            if (container) {
                // Perform operations on the container
                console.log(`Found container: ${name}`, container);
                // Example: Change background color
                container.style.backgroundColor = 'lightblue';
            } else {
                console.log(`Container with name ${name} not found.`);
            }
        }

        // Example of using the function
        // Use this function as needed in your code
        // doSomethingWithContainer('Container 1');