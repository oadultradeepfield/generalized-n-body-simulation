# N-Body Orbit Simulation

![C++](https://img.shields.io/badge/c++-%2300599C.svg?style=for-the-badge&logo=c%2B%2B&logoColor=white)
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Python](https://img.shields.io/badge/python-3670A0?style=for-the-badge&logo=python&logoColor=ffdd54)

This repository builds upon my [previous work on 3-Body Simulation](https://github.com/oadultradeepfield/three-body-simulation/). While the core implementation remains unchanged, this version introduces support for both spherical and Cartesian coordinate systems in a unified configuration file. Note that the angle $\theta$ is measured from the X-axis, while $\phi$ is measured from the Z-axis down to the plane orthogonal to it. Additionally, it calculates changes in velocity resulting from momentum transfer during collisions, assuming elastic collisions.

In January 2025, I reimplemented the entire program in Go, being designed to input and output the same data as the C++ version. The distinct installation steps for each are labeled with (C++) or (Go) before the instructions below.

## Installation (Same as Previous)

- Clone the repository and navigate to the project directory:

  ```bash
  git clone https://github.com/oadultradeepfield/n-body-simulation.git
  cd n-body-simulation
  ```

- (C++) Create a build directory, generate the Makefile with CMake, and build the project:

  ```bash
  mkdir build
  cd build
  cmake ..
  make
  ```

  **Note**: For Go, you can skip this step. Ensure you have Go installed on your device.

- Install Python dependencies for visualization:

  ```bash
  pip install -r python/requirements.txt
  ```

## Usage (Updated Configuration File)

- Configure the simulation by editing the new unified `config.json` file. This file includes all the necessary configuration parameters for both the simulation and the bodies, as well as the coordinate system type. Example:

  ```json
  {
    "config": {
      "G": 6.6743e-11,
      "dt": 1000,
      "total_time": 3.16e7,
      "filename": "results/example_sun_earth_lagrangian_points.txt",
      "collision_distance": 1e-8
    },
    "coordinates_type": "spherical",
    "bodies": [
      {
        "_name": "Sun",
        "mass": 1.989e30,
        "position": [0.0, 0.0, 1.5707963268],
        "velocity": [0.0, 0.0, 0.0]
      },
      {
        "_name": "Earth",
        "mass": 5.972e24,
        "position": [1.496e11, 0.0, 1.5707963268],
        "velocity": [0.0, 2.9788e4, 0.0]
      },
      {
        "_name": "L1",
        "mass": 6500,
        "position": [1.481e11, 0.0, 1.5707963268],
        "velocity": [0.0, 2.9489e4, 0.0]
      }
    ]
  }
  ```

  - `config.json` contains general simulation parameters like the gravitational constant `G`, time step `dt`, total time, output file, and collision distance.
  - The `coordinates_type` specifies the coordinate system (`"cartesian"` or `"spherical"`).
  - The `bodies` section lists the celestial bodies with their respective properties: `name`, `mass`, `position`, and `velocity`.
  - If you are using Cartesian coordinates, make sure to input the coordinates as $x, y, z$. For spherical coordinates, use $r, \theta, \phi$ instead.
  - Note that `_name` is optional; it is simply used to help guide the user when inputting multiple objects, making it easier to navigate through the configuration.

- (C++) Run the simulation, specifying the coordinate system as a parameter (either `cartesian` or `spherical`):

  ```bash
  build/NBodyOrbit config.json
  ```

- (Go) Run the simulation, specifying the coordinate system as a parameter (either `cartesian` or `spherical`):

  ```bash
  go run go/cmd/main.go config.json
  ```

- Generate trajectory plots (optional):

  ```bash
  python3 python/plot.py --filename ./results/example_sun_earth_lagrangian_points.txt --N 7 --labels Sun,Earth,L1,L2,L3,L4,L5
  ```

  The output plot will be saved in the same directory as the `.txt` file.

  |                               **Example 3D Trajectory**                               |                                **Example XY Projection**                                 |
  | :-----------------------------------------------------------------------------------: | :--------------------------------------------------------------------------------------: |
  | ![Example 3D Trajectory](/results/example_sun_earth_lagrangian_points_trajectory.png) | ![Example 2D Projection](/results/example_sun_earth_lagrangian_points_2d_projection.png) |
