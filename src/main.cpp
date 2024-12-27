#include <iostream>
#include <vector>
#include <fstream>
#include <cstdlib>
#include <chrono>
#include <nlohmann/json.hpp>
#include "simulation.h"
#include "bodies.h"
#include "utils.h"
#include "integration.h"

using json = nlohmann::json;

void convert_spherical_to_cartesian(const std::vector<double> &spherical_pos,
                                    const std::vector<double> &spherical_vel,
                                    std::vector<double> &cartesian_pos,
                                    std::vector<double> &cartesian_vel)
{
    double r = spherical_pos[0];
    double theta = spherical_pos[1];
    double phi = spherical_pos[2];

    double vr = spherical_vel[0];
    double vtheta = spherical_vel[1];
    double vphi = spherical_vel[2];

    cartesian_pos = {
        r * std::cos(theta) * std::sin(phi),
        r * std::sin(theta) * std::sin(phi),
        r * std::cos(phi)};

    cartesian_vel = {
        vr * std::sin(phi) * std::cos(theta) - vtheta * std::sin(theta) + vphi * std::cos(theta) * std::cos(phi),
        vr * std::sin(phi) * std::sin(theta) + vtheta * std::cos(theta) + vphi * std::cos(theta) * std::sin(phi),
        vr * std::cos(phi) - vphi * std::sin(phi)};
}

std::vector<Body> parse_bodies_from_json(const json &j)
{
    std::vector<Body> bodies;
    std::string coordinates_type = j["coordinates_type"];

    for (const auto &body : j["bodies"])
    {
        double mass = body["mass"];
        std::vector<double> pos = body["position"];
        std::vector<double> vel = body["velocity"];

        if (coordinates_type == "spherical")
        {
            std::vector<double> cartesian_pos, cartesian_vel;
            convert_spherical_to_cartesian(pos, vel, cartesian_pos, cartesian_vel);
            bodies.push_back(Body(mass, cartesian_pos, cartesian_vel));
        }
        else
        {
            bodies.push_back(Body(mass, pos, vel));
        }
    }

    return bodies;
}

int main(int argc, char *argv[])
{
    if (argc < 2)
    {
        std::cerr << "Usage: " << argv[0] << " <json_config_file>" << std::endl;
        return 1;
    }

    std::string json_filename = argv[1];

    std::ifstream file(json_filename);
    if (!file)
    {
        std::cerr << "Could not open JSON file: " << json_filename << std::endl;
        return 1;
    }

    json j;
    try
    {
        file >> j;
    }
    catch (const json::parse_error &e)
    {
        std::cerr << "JSON parsing error: " << e.what() << std::endl;
        return 1;
    }

    double G = j["config"]["G"];
    double dt = j["config"]["dt"];
    double total_time = j["config"]["total_time"];
    std::string output_filename = j["config"]["filename"];
    double collision_distance = j["config"]["collision_distance"];

    std::vector<Body> bodies = parse_bodies_from_json(j);

    std::cout << "Simulation started with " << bodies.size() << " bodies..." << std::endl;

    auto start_time = std::chrono::high_resolution_clock::now();

    run_simulation(bodies, dt, total_time, output_filename, G, collision_distance);

    auto end_time = std::chrono::high_resolution_clock::now();
    std::chrono::duration<double> elapsed_time = end_time - start_time;

    std::cout << "Simulation completed. Results saved to " << output_filename << std::endl;
    std::cout << "Time taken: " << elapsed_time.count() << " seconds" << std::endl;

    return 0;
}