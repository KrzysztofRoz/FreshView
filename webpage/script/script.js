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
                console.log(data);
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
                displayDuties(duties);
                });
            }

        // Function to display duties
        function displayDuties(duties) {
            const dutiesDisplay = document.querySelector('.dutiesDisplay');
            const dutiesList = document.getElementById('dutiesList');
            dutiesList.innerHTML = ''; // Clear previous duties

            duties.forEach(duty => {
                const dutyDiv = document.createElement('div');
                dutyDiv.className = 'duty';

                const taskNameDiv = document.createElement('div');
                taskNameDiv.className = 'taskName';
                taskNameDiv.textContent = `${duty.TaskName}`;

                const categoryDiv = document.createElement('div');
                categoryDiv.className = 'category';
                categoryDiv.textContent = `Category: ${duty.Category}`;

                const createdAtDiv = document.createElement('div');
                createdAtDiv.className = 'createdAt';
                createdAtDiv.textContent = `Created At: ${new Date(duty.CreatedAt).toLocaleString()}`;

                dutyDiv.appendChild(taskNameDiv);
                dutyDiv.appendChild(categoryDiv);
                dutyDiv.appendChild(createdAtDiv);

                dutiesList.appendChild(dutyDiv);
            });

            dutiesDisplay.classList.add('active');
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