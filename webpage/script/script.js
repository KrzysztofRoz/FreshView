        // Function to fetch containers from the backend
        async function fetchContainers() {
            try {
                const response = await fetch('http://localhost:8080/v1/api/retreive/containers/all', {
                    method: 'GET',
                    headers: {
                        
                        'Content-Type': 'application/json',
                        'FreshView-API-Key':'apiKey'
                    },

                });
                const data = await response.json();
                console.log(data)
                return data.containerNames; 
            } catch (error) {
                console.error('Error fetching containers:', error);
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
            dutieContainerDiv.addEventListener('click',() => {
                document.querySelector('.main').classList.toggle('shrink');
            })
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