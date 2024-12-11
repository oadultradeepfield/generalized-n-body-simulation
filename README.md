# Generalized N-Body Simulation

This repository builds upon my [previous work on 3-Body Simulation](https://github.com/oadultradeepfield/three-body-simulation/). While the core implementation remains unchanged, this version introduces support for both spherical and Cartesian coordinate systems. Note that the angle $\theta$ is measured from the X-axis, while $\phi$ is measured from the Z-axis down to the plane orthogonal to it.

## Installation (Same as Previous)

1. Clone the repository and navigate to the project directory:

   ```bash
   git clone https://github.com/oadultradeepfield/three-body-simulation.git
   cd three-body-simulation
   ```

2. Create a build directory, generate the Makefile with CMake, and build the project:

   ```bash
   mkdir build
   cd build
   cmake ..
   make
   ```

3. Install Python dependencies for visualization:

   ```bash
   pip install -r python/requirements.txt
   ```

## Usage (Extra Argument to Specify Coordinate System)

1. Configure the simulation by editing `config.txt` and `bodies_cartesian.txt` or `bodies_spherical.txt`:

   - `config.txt`

     ```bash
        G=6.6743e-11
        dt=1000
        total_time=3.16e7
        filename=results/example_sun_earth_lagrangian_points.txt
     ```

   - `bodies_spherical.txt`

     ```bash
     # Sun
     1.989e30
     0.0 0.0 1.5707963268
     0.0 0.0 0.0

     # Earth
     5.972e24
     1.496e11 0.0 1.5707963268
     0.0 2.9788e4 0.0

     # L1
     6.500e3
     1.481e11 0.0 1.5707963268
     0.0 2.9489e4 0.0
     ...
     ```

2. Run the simulation (change `spherical` to `cartesian` if you wish to do so):

   ```bash
   build/NBodySimulation config.txt bodies_spherical.txt spherical
   ```

3. Generate trajectory plots (optional):

   ```bash
   python3 python/plot.py --filename ./results/example_sun_earth_lagrangian_points.txt --N 7 --labels Sun,Earth,L1,L2,L3,L4,L5
   ```

The output plot will be saved in the same directory as the `.txt` file.

|                               **Example 3D Trajectory**                               |                                **Example XY Projection**                                 |
| :-----------------------------------------------------------------------------------: | :--------------------------------------------------------------------------------------: |
| ![Example 3D Trajectory](/results/example_sun_earth_lagrangian_points_trajectory.png) | ![Example 2D Projection](/results/example_sun_earth_lagrangian_points_2d_projection.png) |
